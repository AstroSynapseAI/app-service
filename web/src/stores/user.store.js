import { defineStore } from 'pinia';
import { fetchWrapper } from '../helpers/fetch-wrapper.js';
import { useChatStore } from './chat.store.js';
import { useStorage } from '@vueuse/core';

const apiUrl = import.meta.env.VITE_API_URL;
const usersURL = `${apiUrl}/users`;

export const useUserStore = defineStore({
  id: 'users',
  state: () => ({
    current: useStorage('user', {}),
    avatar: useStorage('avatar', {}),
    session_id: useStorage('session_id', null),
    account: {},
    record: {},
    records: [],
  }),
  getters: {
  },
  actions: {
    async get() {
      try {
        const currentUser = await fetchWrapper.get(`${usersURL}/${this.current.ID}`);
        this.current = currentUser
      } catch (error) {
        console.error(error);
      }
    },
    
    async hasAvatar(user_id) {
      try {
        const avatar = await fetchWrapper.get(`${usersURL}/${user_id}/avatars`);
        this.avatar = avatar
        return true
      } catch (error) {
        console.error(error);
        if (error.status === 404) {
          return false
        }
      }
    },
    async getSessionToken() {
      try {
        const session = await fetchWrapper.get(`${apiUrl}/users/token`);
        this.session_id = session.token;
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
        this.avatar = avatar
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
        throw error.Error
      }
  },
  }
  
})
