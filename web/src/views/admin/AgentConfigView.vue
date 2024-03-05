<script setup>
import { Form, Field } from 'vee-validate';
import { ref, toRef, onMounted } from 'vue';
import { useRoute } from 'vue-router';
import { useAvatarStore } from '@/stores/avatar.store';
import { useUserStore } from '@/stores/user.store';
import { useAgentStore } from '@/stores/agent.store';
import { useLLMStore } from '@/stores/llm.store';
import { useToolStore } from '@/stores/tool.store';

const user = useUserStore();
const avatar = useAvatarStore();
const route = useRoute();

const agent = useAgentStore();
const agentRecord = toRef(agent, 'record');
const activeTools = toRef(agent, 'activeTools');

const llm = useLLMStore();
const llmRecords = toRef(llm, 'records');

const tool = useToolStore();
const toolRecords = toRef(tool, 'records');

const agentPrimer = ref('');
const agentName = ref('');
const activeAgentLLMID = ref('');
const isActiveAgent = ref(false);
const isPublicAgent = ref(false);

const toolIsActive = (toolID) => {
  const activeTool = activeTools.value.find(activeTool => {
    return activeTool.tool.ID == toolID
  });

  return activeTool ? activeTool.is_active : false
}

const toggleActive = () => {
  isActiveAgent.value = !isActiveAgent.value
}

const togglePublic = () => {
  isPublicAgent.value = !isPublicAgent.value
}

const toggleTool = (tool) => {
  const activeTool = activeTools.value.find( activeTool => {
    return activeTool.tool.ID == tool.ID;
  });

  if (!activeTool) {
    const toolData = {
      active_agent_id: agentRecord.value.ID,
      tool_id: tool.ID,
      token: tool.token,
      tool: tool,
      is_active: true,
    }
    activeTools.value.push(toolData);
    return
  }

  activeTool.is_active = !activeTool.is_active;
  activeTool.token = tool.token;

  console.log("activeTools", activeTools.value);
}

const submitForm = async () => {
  const formData = {
    ID: parseInt(route.params.active_agent_id),
    avatar_id: parseInt(route.params.avatar_id),
    agent_id: parseInt(route.params.agent_id),
    llm_id: activeAgentLLMID.value,
    primer: agentPrimer.value,
    is_active: isActiveAgent.value,
    is_public: isPublicAgent.value,
  }

  if (agentRecord.value.ID) {
    formData.ID = agentRecord.value.ID;
  }

  try {
    await agent.saveActiveAgent(formData);
    for (const activeTool of activeTools.value) {
      await tool.saveAgentTool(activeTool);
    }
  }
  catch (error) {
    console.log(error);
  }
}

onMounted(async () => {
  try {
    await llm.getLLMs();
    await tool.getTools();
    await agent.getAgent(route.params.agent_id);
    if (agentRecord.value) {
      agentPrimer.value = agentRecord.value.primer;
      agentName.value = agentRecord.value.name;  
    }
  
    if (route.params.active_agent_id) {
      await avatar.getActiveAgent(route.params.agent_id, route.params.avatar_id);
      await agent.getActiveTools(avatar.activeAgent.agent_id);
      if (avatar.activeAgent) {
        activeAgentLLMID.value = avatar.activeAgent.llm_id;
        agentPrimer.value = avatar.activeAgent.primer;
        isActiveAgent.value = avatar.activeAgent.is_active;
        isPublicAgent.value = avatar.activeAgent.is_public;
      }
    }

    for (const [index, tool] of toolRecords.value.entries()) {
      const activeTool = activeTools.value.find(activeTool => {
        return activeTool.tool.ID == tool.ID;
      });

      if (activeTool) {
        tool.token = activeTool.token;
      }

      toolRecords.value[index] = tool;
    }

    console.log("loaded tool records:", toolRecords.value);
  }
  catch (error) {
    console.log(error);
  }
  
  feather.replace();
});

</script>


<template>

  <div class="container-fluid p-0">
    <h1 class="h3 mb-3">Configure: {{ agentRecord.name }}
      <div class="form-check form-switch float-end me-5">
        <label class="form-check-label" for="flexSwitchCheckDefault">Active</label>
        <input class="form-check-input" type="checkbox" id="flexSwitchCheckDefault" :checked="isActiveAgent" @click="toggleActive">
      </div>
      <div class="form-check form-switch float-end me-3">
        <label class="form-check-label" for="flexSwitchCheckDefault">Public</label>
        <input class="form-check-input" type="checkbox" id="flexSwitchCheckDefault" :checked="isPublicAgent" @click="togglePublic">
      </div>
    </h1>

    <div class="row">
      <div class="col-12">

        <div class="card">

          <div class="card-body">
            <div class="container">

              <div class="row">
                <div class="col-12">
                  <h3>Primer</h3>
                  <Field
                    v-model="agentPrimer"
                    name="agent_primer"
                    type="text"
                    as="textarea"
                    class="form-control"
                    rows="8"
                    placeholder="How should the agent behave..."
                  ></Field>
                </div>
              </div>

              <div class="row mt-3">
                <div class="col-12">
                  <h3>Agent Model</h3>
                  <select v-model="activeAgentLLMID" class="form-select model-select" aria-label="Select Model">
                    <option value="" disabled selected>Select a LLM</option>
                    <option v-for="(llm, index) in llmRecords" :value="llm.ID" :key="index">
                      {{ llm.name }}
                    </option>
                  </select>
                </div>
              </div>

              <div class="row mt-3">
                <div class="col-12" v-if="agentRecord.slug === 'search-agent'">
                  <h3>Agent Tools</h3>
                  <ul class="list-group">
                    <li v-for="(tool, index) in toolRecords" :key="index" class="list-group-item">
                      <span>{{ tool.name }}</span>
                      
                      <div class="form-floating mb-3" v-if="tool.slug != 'ddg-search'">
                        <input type="text" class="form-control" id="apiToken" placeholder="Tool Token" v-model="tool.token">
                        <label for="apiToken">Api Token</label>
                      </div>
                      
                      <i data-feather="settings" class="float-end"></i>
                      <div class="form-check form-switch float-end me-3">
                        <label class="form-check-label" for="checkboxTool{{ tool.ID }}">Active</label>
                        <input class="form-check-input" type="checkbox" id="checkboxTool{{ tool.ID }}" @click="toggleTool(tool)" :checked="toolIsActive(tool.ID)">
                      </div>
                    </li>
                  </ul>

                </div>
              </div>
              
              <div class="row mt-3">
                <div class="col-12">
                  <button type="button" class="btn btn-secondary" @click="submitForm">Save</button>
                </div>    
              </div>

            </div>

          </div>

        </div>

      </div>
    </div>
  </div>

  </template>
