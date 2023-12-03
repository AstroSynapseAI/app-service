import { defineStore } from 'pinia';

export const useAgentsStore = defineStore({
  id: 'agents',
  state: () => ({
    agents: [],
    agent: {},
    activeAgents: [],
    activeAgent: {},
  }),
  actions: {
    async getAgents() {
      
    }, 

    async getAgent() {
      
    },

    async getAgnetTools() {
      
    }, 

    async updateAgent() {
      
    },

    async setActiveAgent() {
      
    },

    async setInactiveAgent() {
      
    },

    async publishAgent() {
      
    },

    async unpublishAgent() {
      
    }, 

    async setAgentModel() {
      
    },

    async setActiveTool() {
      
    }, 

    async setInactiveTool() {
      
    }

  }
})