import { defineStore } from 'pinia';

export const usePluginsStore = defineStore({
  id: 'plugins',
  state: () => ({
    plugins: [],
    plugin: {},
  }),
  actions: {
    async getPlugins() {
      
    },

    async getPlugin() {
      
    },

    async updatePlugin() {
      
    },

    async deletePlugin() {
      
    }
  }
})