import { defineStore } from 'pinia';
import { fetchWrapper } from '../helpers/fetch-wrapper.js';
import { useChatStore } from './chat.store.js';

const apiUrl = import.meta.env.VITE_API_URL;
const usersURL = `${apiUrl}/users`;

export const useUserStore = defineStore({
  id: 'users',
  state: () => ({
    account: {},
    record: {},
    records: [],
  }),
  getters: {
    current: () => JSON.parse(localStorage.getItem('user')),
    avatar: () => JSON.parse(localStorage.getItem('avatar')),
    session_id: () => localStorage.getItem('session_id'),
  },
  actions: {
    async getSessionToken() {
      try {
        const session = await fetchWrapper.get(`${apiUrl}/users/token`);
        localStorage.setItem('session_id', session.token);
      } catch (error) {
        console.error(error);
      }
    },

    async getUsers() {
      try {
        const users = await fetchWrapper.get(`${usersURL}`);
        this.records = users;
      } catch (error) {
        console.error(error);
      }
    },

    async getUserAvatar(user_id) {
      try {
        const avatar = await fetchWrapper.get(`${usersURL}/${user_id}/avatars`);
        localStorage.setItem('avatar', JSON.stringify(avatar));
      } catch (error) {
        console.error(error);
      }
    },

    async getUserAccounts(user_id) {
      try {
        const accounts = await fetchWrapper.get(`${usersURL}/${user_id}/accounts`);
        this.account = accounts[0];
      } catch (error) {
        console.error(error);
      }
    },

    // legacy function
    async getSession() {
      console.log("Creating session...")
      const chatStore = useChatStore();
      this.user = JSON.parse(localStorage.getItem('user'));

      if (!this.user) {
        try {
          chatStore.messages = [{
              sender: "ai",
              content: "Hello there... I'm Asai, How can I help you?"
          }];
          const user = await fetchWrapper.get(`${usersURL}/session`);
          localStorage.setItem('user', JSON.stringify(user));
        } catch (error) {
          console.error(error);
        }
      }

      if (this.user) {
        console.log("User:", this.user);
        chatStore.loadHistory();
      }
    }, 

    async saveProfile(user_id, formData) {
      try {
        const user = await fetchWrapper.post(`${usersURL}/${user_id}/save/profile`, formData);
        localStorage.setItem('user', JSON.stringify(user));
      } catch (error) {
        throw error.Error // Rethrow the error to be handled in the component
      }
    }, 

    async changeEmail(user_id, formData) {
      try {
        await fetchWrapper.put(`${usersURL}/${user_id}/change/email`, formData);
      } catch (error) {
        console.error(error);
      }
    },

    async changePassword(user_id, formData) {
      try {
        await fetchWrapper.put(`${usersURL}/${user_id}/change/password`, formData);
      } catch (error) {
        console.error(error);
      }
    }
  }

})