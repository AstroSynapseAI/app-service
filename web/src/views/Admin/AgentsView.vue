<script setup>
  import { useUsersStore } from '../stores/user.store';

  const userStore = useUsersStore();
  const avatarName = userStore.user.avatar
</script>

<template>
          
  <div class="container-fluid p-0">
    <h1 class="h3 mb-3">Agents</h1>
    <div class="row">
      <div class="col-12">
        <div class="container">
          
          <div class="row">
            <div class="col-6">
              <div class="card">
                <div class="card-header">
                  <div class="row">
                    <div class="col-6">
                      <h5 class="card-title">Search Agent</h5>
                    </div>
                    <div class="col-6">
                      <div class="form-check form-switch float-end">
                        <label class="form-check-label" for="flexSwitchCheckDefault">Active</label>
                        <input class="form-check-input" type="checkbox" id="flexSwitchCheckDefault">
                      </div>
                    </div>
                  </div>

                  
                </div>
                <div class="card-body">
                  <p>Utilizes search engines such as Google, DuckDuckGo, and Metaphor for automated web searches.</p>
                  <div>
                    <router-link :to="{name: 'agent-config', params: { avatar: avatarName, slug: 'search' }}" class="btn btn-primary">Configure</router-link>
                  </div>
                </div>
              </div>
            </div>
            <div class="col-6">
              <div class="card">
                <div class="card-header">

                  <div class="row">
                    <div class="col-6">
                      <h5 class="card-title">Browser Agent</h5>
                    </div>
                    <div class="col-6">
                      <div class="form-check form-switch float-end">
                        <input class="form-check-input" type="checkbox" id="flexSwitchCheckDefault">
                        <label class="form-check-label" for="flexSwitchCheckDefault">Active</label>
                      </div>
                    </div>
                  </div>

                </div>
                <div class="card-body">
                  <p>Equipped with the capability to scrape, read website contents, and interact with web pages and web applications.</p>
                  <div>
                    <router-link :to="{name: 'agent-config', params: { avatar: avatarName, slug: 'browser' }}" class="btn btn-primary">Configure</router-link>
                  </div>
                </div>
              </div>
            </div>
          </div>


        </div>
      </div>
    </div>
  </div>
      
</template>