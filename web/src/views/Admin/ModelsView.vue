<script setup>

import { onMounted } from 'vue';

onMounted(() => {
  feather.replace();
});
</script>

<template>
          
  <div class="container-fluid p-0">
    <h1 class="h3 mb-3">Models
      <button class="btn btn-primary float-end">Add Models <i data-feather="plus-square"></i></button>
    </h1>
    <div class="card">

      <div class="card-body">
      <div class="container p-4">
        <ul class="list-group">
        
          <li class="list-group-item">
            
            <div class="row">
              <div class="col-4"><span>LLama2</span></div>
              <div class="col-6"><span class="badge bg-success">Local</span></div>
              <div class="col-2">
                <i data-feather="settings" class="float-end"></i>
                <div class="form-check form-switch float-end me-3">
                  <label class="form-check-label" for="flexSwitchCheckDefault">Active</label>
                  <input class="form-check-input" type="checkbox" id="flexSwitchCheckDefault">
                </div>
              </div>
            </div>

          </li>
        
          <li class="list-group-item">
            
            <div class="row">
              <div class="col-4"><span>Falcon LLM</span></div>
              <div class="col-6"><span class="badge bg-success">Local</span></div>
              <div class="col-2">
                <i data-feather="settings" class="float-end"></i>
                <div class="form-check form-switch float-end me-3">
                  <label class="form-check-label" for="flexSwitchCheckDefault">Active</label>
                  <input class="form-check-input" type="checkbox" id="flexSwitchCheckDefault">
                </div>
              </div>
            </div>
          
          </li>
        
          <li class="list-group-item">
          
            <div class="row">
              <div class="col-4"><span>Dolphin LLM</span></div>
              <div class="col-6"><span class="badge bg-success">Local</span></div>
              <div class="col-2">
                <i data-feather="settings" class="float-end"></i>
                <div class="form-check form-switch float-end me-3">
                  <label class="form-check-label" for="flexSwitchCheckDefault">Active</label>
                  <input class="form-check-input" type="checkbox" id="flexSwitchCheckDefault">
                </div>
              </div>
            </div>

          </li>

          <li class="list-group-item">
            
            <div class="row">
              <div class="col-4"><span>GPT-4</span></div>
              <div class="col-6">
                <span class="badge bg-primary me-3">API</span>
                <span class="badge bg-danger">Additional costs apply!</span>
              </div>
              <div class="col-2">
                <i data-feather="settings" class="float-end"></i>
                <div class="form-check form-switch float-end me-3">
                  <label class="form-check-label" for="flexSwitchCheckDefault">Active</label>
                  <input class="form-check-input" type="checkbox" id="flexSwitchCheckDefault">
                </div>
              </div>
            </div>

          </li>

          <li class="list-group-item">
            
            <div class="row">
              <div class="col-4"><span>Claude</span></div>
              <div class="col-6">
                <span class="badge bg-primary me-3">API</span>
                <span class="badge bg-danger">Additional costs apply!</span>
              </div>
              <div class="col-2">
                <i data-feather="settings" class="float-end"></i>
                <div class="form-check form-switch float-end me-3">
                  <label class="form-check-label" for="flexSwitchCheckDefault">Active</label>
                  <input class="form-check-input" type="checkbox" id="flexSwitchCheckDefault">
                </div>
              </div>
            </div>

          </li>
          
        </ul>
    
        <div class="row"></div>
      </div>
    </div>
  </div>

  </div>
        
  </template>