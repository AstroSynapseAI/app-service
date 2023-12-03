import { defineStore } from 'pinia';

export const usePersonasStore = defineStore({
  id: 'personas',
  state: () => ({
    personas: [],
    persona: {},
  }),
  actions: {
    async getPersonas() {
      
    },

    async getPersona() {
      
    },

    async createPersona() {
      
    },

    async updatePersona() {
      
    },

    async deletePersona() {
      
    }
  }
})