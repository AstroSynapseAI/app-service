<script setup>
import { Form, Field } from 'vee-validate';
import { ref, onMounted } from 'vue';
import { useRoute } from 'vue-router';

const route = useRoute();
let agentName = ref(route.params.slug);

onMounted(() => {
  feather.replace();
});

</script>


<template>
          
  <div class="container-fluid p-0">
    <h1 class="h3 mb-3">Configure: {{ agentName }}
      <div class="form-check form-switch float-end me-5">
        <label class="form-check-label" for="flexSwitchCheckDefault">Active</label>
        <input class="form-check-input" type="checkbox" id="flexSwitchCheckDefault">
      </div>
      <div class="form-check form-switch float-end me-3">
        <label class="form-check-label" for="flexSwitchCheckDefault">Public</label>
        <input class="form-check-input" type="checkbox" id="flexSwitchCheckDefault">
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
                  <select class="form-select model-select" aria-label="Select Model">
                    <option selected>GPT-4</option>
                    <option value="3">Claude LLM</option>
                    <option value="1">LLama2</option>
                    <option value="2">Falcon LLM</option>
                  </select>
                </div>
              </div>

              <div class="row mt-3">
                <div class="col-12">
                  <h3>Agent Tools</h3>
                  <ul class="list-group">
                    <li class="list-group-item">
                      <span>Google Search via Serpa API</span>
                      <i data-feather="settings" class="float-end"></i>
                      <div class="form-check form-switch float-end me-3">
                        <label class="form-check-label" for="flexSwitchCheckDefault">Active</label>
                        <input class="form-check-input" type="checkbox" id="flexSwitchCheckDefault">
                      </div>
                    </li>
                    <li class="list-group-item">
                      <span>DuckDuckGo Search API</span>
                      <i data-feather="settings" class="float-end"></i>
                      <div class="form-check form-switch float-end me-3">
                        <label class="form-check-label" for="flexSwitchCheckDefault">Active</label>
                        <input class="form-check-input" type="checkbox" id="flexSwitchCheckDefault">
                      </div>
                    </li>
                    <li class="list-group-item">
                      <span>Metaphor Search API</span>
                      <i data-feather="settings" class="float-end"></i>
                      <div class="form-check form-switch float-end me-3">
                        <label class="form-check-label" for="flexSwitchCheckDefault">Active</label>
                        <input class="form-check-input" type="checkbox" id="flexSwitchCheckDefault">
                      </div>
                    </li>
                    
                  </ul>

                </div>
              </div>



            </div>

          </div>
        
        </div>
      
      </div>
    </div>
  </div>
        
  </template>