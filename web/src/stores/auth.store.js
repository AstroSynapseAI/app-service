import { defineStore } from 'pinia'
import { fetchWrapper }  from '../helpers/fetch-wrapper.js';
import router from '@/router/index.js';

const usersURL = `${import.meta.env.VITE_API_URL}/users`;

export const useAuthStore = defineStore({
  id: 'auth',
  state: () => ({
    currentUser: JSON.parse(localStorage.getItem('user')),
    isLogedIn: false,
    apiToken: null
  }),
  actions: {
   async login(username, password) {
      const reqBody = {
        username: username,
        password: password
      }
      try {
        const user = await fetchWrapper.post(`${usersURL}/login`, reqBody);
        if (user) {
          if (user.apiToken) {
            this.apiToken = user.apiToken
          }
          this.isLogedIn = true
          localStorage.setItem('user', JSON.stringify(user));
          return true
        }
        return false
      } catch (error) {
        console.error(error);
        localStorage.removeItem('user');
        this.isLogedIn = false
        this.apiToken = null
        return false
      }
    },

    async inviteUser(username) {
      const reqBody = {
        username: username
      }
      try {
        const user = await fetchWrapper.post(`${usersURL}/invite`, reqBody);
        return user
      } catch (error) {
        console.error(error);
        return false
      }
    },

    async registerinvite(formData) {
      try {
        const user = await fetchWrapper.post(`${usersURL}/register/invite`, formData);
        if (user) {
          if (user.apiToken) {
            this.apiToken = user.apiToken
          }
          this.isLogedIn = true
          localStorage.setItem('user', JSON.stringify(user));
          return true
        }
        return false
      } catch (error) {
        console.error(error);
        return false
      }
    },

    logout() {
      localStorage.removeItem('user');
      this.isLogedIn = false
      this.apiToken = null
      router.push('/login');
    }
  }
})