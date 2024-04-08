import { defineStore } from 'pinia';
import { fetchWrapper }  from '../helpers/fetch-wrapper.js';
import { useUserStore } from './user.store.js';

const chatURL = `${import.meta.env.VITE_API_URL}/chat`;
const wsURL = `${import.meta.env.VITE_WS_URL}/ws/chat`;
const apiUrl = import.meta.env.VITE_API_URL;

export const useChatStore = defineStore({
  id: 'chat',
  state: () => ({
    socket: null,
    retryCount: 0,
    maxRetryCount: 3,
    isLoading: false,
    messages: [],
    // messages: [{
    //   sender: "ai",
    //   content: "Hello there... I'm Asai, How can I help you?",
    //   isLoading: false,
    //   isAgentRunnig: false,
    //   agentName: null
    // }],
    aiMsg: {
      sender: "ai",
      content: "",
      isLoading: false,
      isAgentRunnig: false,
      agentName: null
    },
    connectionErr: {
      active: false,
      msg: null
    }
  }),
  actions: {
    async connectWebSocket() {
      this.socket = new WebSocket(wsURL);

      this.socket.addEventListener('open', (event) => {
        this.connectionErr.status = false;
      })

      this.socket.addEventListener('error', (event) => {
        console.log("Connection error:", event);
        this.connectionErr.status = true;
      });

      this.socket.addEventListener('close', (event) => {
        this.retryWSConnection()
      });

      this.socket.addEventListener('message', (event) => {
        var payload = JSON.parse(event.data);

        switch (payload.step) {
          case "chain start":
            this.onChainStart(payload);
            break;
          case "chain end":
            this.onChainEnd(payload);
            break;
          case "agent action":
            this.onAgentRunning(payload);
            break;
          case "final output":
            this.onMessage(payload);
            break;
          case "error":
            this.connectionErr.active = true;
        }
      });
    },
    onChainStart(payload) {
      this.aiMsg = {
        sender: "ai",
        content: "",
        isLoading: true,
        isAgentRunnig: false,
        agentName: null
      }

      this.messages = [...this.messages, this.aiMsg];
    },
    onChainEnd(payload) {
      this.isLoading = false
      this.aiMsg.isLoading = false;
      this.aiMsg.isAgentRunnig = false;
      this.aiMsg.agentName = null;
    },
    onAgentRunning(payload) {
      this.aiMsg.isAgentRunnig = true;
      this.aiMsg.agentName = payload.agent;
    },
    onMessage(payload) {
      if (this.aiMsg.isLoading) {
        this.aiMsg.isLoading = false;
      }
      this.aiMsg.content += payload.msg;
    },
    retryWSConnection() {
      console.log("Retrying connection...");
      if(!navigator.onLine) {
        window.addEventListener('online', this.reconnect);
      } else {
        this.reconnect();
      }
    },
    reconnect() {
      if (this.retryCount < this.maxRetryCount) {
        this.retryCount++;
        setTimeout(() => this.connectWebSocket(), 3000);
      } else {
        console.log('Failed to reconnect to the WebSocket after '+this.maxRetryCount+' attempts.');
      }
    },

    async sendPrompt(payload) {      
      console.log("Sending prompt:", payload);

      var userMsg = {
        sender: "human",
        content: payload.msg
      }

      this.messages = [...this.messages, userMsg];
      
      const reqData = {
        session_id: payload.session_id,
        avatar_id: payload.avatar_id,
        user_prompt: payload.msg
      }

      try {
        this.socket.send(JSON.stringify(reqData));
        this.isLoading = true;
      } catch (error) {
        console.error("Failed to send prompt:", error);
      }
    },

    async getPrivateSession(avatar_id, session_id) {
      console.log("Loading session...");
      var responseMsgs = []
      try {
        const histroy = await fetchWrapper.get(`${apiUrl}/avatars/${avatar_id}/session/${session_id}`);
        if (histroy) {
          for (var i = 0; i < histroy.length; i++) {
            var msg = {
              sender: histroy[i].type,
              content: histroy[i].text
            }
            responseMsgs.push(msg)
          }
          this.messages = responseMsgs;
        }
      } catch (error) {
        console.error(error);
      }
    },

    closeError() {
      this.connectionErr.active = false;
    },

    newUserConnected(userStore) {
      const reqData = {
        session_id: userStore.user.session_id,
        user_prompt: "new user connected"
      }

      try {
        this.socket.send(JSON.stringify(reqData));
        this.isLoading = true;
      } catch (error) {
        console.error("Failed to send prompt:", error);
      }
    }
  }
})
