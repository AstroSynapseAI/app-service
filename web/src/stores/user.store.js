import { defineStore } from 'pinia';

import { fetchWrapper } from '../helpers/fetch-wrapper.js';
import { useChatStore } from './chat.store.js';

const apiUrl = import.meta.env.VITE_API_URL;
const usersURL = `${apiUrl}/users`;

export const useUsersStore = defineStore({
  id: 'users',
  state: () => ({
    user: JSON.parse(localStorage.getItem('user')),
  }),
  actions: {
    async getSession() {
      console.log("Creating session...")
      const chatStore = useChatStore(); 
      this.user = JSON.parse(localStorage.getItem('user'));
      
      if (!this.user) {
        try {
          const user = await fetchWrapper.get(`${usersURL}/session`);
          localStorage.setItem('user', JSON.stringify(user));
          this.user = user;
        } catch (error) {
          console.error(error);
        }    
      }

      if (this.user) {
        chatStore.loadHistory();
      }
    },
    async getUsers() {

    },

    async getUser() {

    },

    async getUserAvatar() {

    },

    async getUserAgents() {

    },

    async getUserPlugins() {

    },

    async getUserModels() {

    },
  }

})