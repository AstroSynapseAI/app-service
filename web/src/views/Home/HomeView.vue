<script setup>
import { onMounted } from 'vue';


onMounted(() => {
  feather.replace();
});
</script>

<template>
  <div class="container d-flex flex-column vh-100">
    
    <!-- Navigation  -->

    <nav class="navbar navbar-expand-md bg-dark bg-transparent">
      <div class="container-fluid">
        <div class="row w-100">
          <div class="col-xl-6 col-md-4">
            <span class="navbar-brand text-white">AI Avatars Platform</span>
          </div>
          <div class="navbar-menu col-xl-6 col-md-8 d-flex justify-content-md-end justify-content-sm-start">

            <div class="row">
              <div class="col-auto">
                <router-link :to="{name: 'asai'}" target="_blank" class="btn text-white px-md-4 px-lg-5 px-sm-2">
                  <!-- <img src="@/assets/ASAILogomark.svg" height="24" alt=""> -->

                  <svg class="asai-icon" height="24" viewBox="0 0 165 247" fill="none" xmlns="http://www.w3.org/2000/svg">
                    <g clip-path="url(#clip0_1100_3274)">
                    <line x1="161.602" y1="1.28433e-07" x2="161.602" y2="246.81" stroke="white" stroke-width="5.87643"/>
                    <line x1="2.93821" y1="1.28433e-07" x2="2.9382" y2="246.81" stroke="white" stroke-width="5.87643"/>
                    <line x1="164.54" y1="245.341" x2="0.000106555" y2="245.341" stroke="white" stroke-width="5.87643"/>
                    <line x1="126.343" y1="155.725" x2="38.1969" y2="155.725" stroke="white" stroke-width="5.87643"/>
                    </g>
                    <defs>
                    <clipPath id="clip0_1100_3274">
                    <rect width="164.54" height="246.81" fill="white" transform="translate(164.54 246.81) rotate(-180)"/>
                    </clipPath>
                    </defs>
                  </svg>

                  <span class="ms-3">Talk to Asai</span>
                </router-link>
              </div>
              <div class="col-auto">
                <router-link :to="{name: 'about', hash: '#content'}" class="btn text-white px-md-5 px-sm-0">
                  <span>About</span>
                </router-link>
              </div>

              <div class="col-auto">
                <router-link :to="{name: 'login'}" class="btn text-white">
                  <i class="align-middle" data-feather="log-in"></i>
                </router-link>
              </div>
            </div>

          </div>
        </div>
      </div>
    </nav>

    <!-- Navigation End -->

    <hr class="top-hr opacity-100">

    <!-- Content -->
    
    <div class="sub-menu row">
      <div class="col-lg-6">
        <img class="logo" src="@/assets/ASAILogotype.svg" alt="">
      </div>

      <div class="col-lg-6">
        <div class="row">

          <div class="col-md-6">
            <table>
              <tr>
                <td>
                  <router-link :to="{name: 'about', hash: '#content', params: {slug: 'open-source'}}">
                    <h5 class="mb-3">Open source</h5>
                  </router-link>
                </td>
                <td class="d-flex justify-content-end">
                  <i class="mb-3 ms-3" data-feather="arrow-right"></i>
                </td>
              </tr>
              <tr>
                <td>
                  <router-link :to="{name: 'about', hash: '#content', params: {slug: 'transparency'}}">
                    <h5 class="mb-3">Transparent</h5>
                  </router-link>
                </td>
                <td class="d-flex justify-content-end">
                  <i class="mb-3" data-feather="arrow-right"></i>
                </td>
              </tr>

              <tr>
                <td>
                  <router-link :to="{name: 'about', hash: '#content', params: {slug: 'personalization'}}">
                    <h5 class="mb-3">Personalized</h5>
                  </router-link>
                </td>
                <td class="d-flex justify-content-end">
                  <i class="mb-3" data-feather="arrow-right"></i>
                </td>
              </tr>

              <tr>
                <td>
                  <router-link :to="{name: 'about', hash: '#content', params: {slug: 'privacy'}}">
                    <h5>Private</h5>
                  </router-link>
                </td>
                <td class="d-flex justify-content-end">
                  <i class="mb-3" data-feather="arrow-right"></i>
                </td>
              </tr>
            </table>
          </div>

          <div class="col-md-6">
            <table>
              <tr>
                <td>
                  <router-link :to="{name: 'about', hash: '#content', params: {slug: 'cloud'}}">
                    <h5 class="mb-3">Cloud</h5>
                  </router-link>
                </td>
                <td class="d-flex justify-content-end">
                  <i class="mb-3 ms-3" data-feather="arrow-right"></i>
                </td>
              </tr>
              <tr>
                <td>
                  <router-link :to="{name: 'about', hash: '#content', params: {slug: 'self-hosted'}}">
                    <h5 class="mb-3">Self-Hosted</h5>
                  </router-link>
                </td>
                <td class="d-flex justify-content-end">
                  <i class="mb-3" data-feather="arrow-right"></i>
                </td>
              </tr>

              <tr>
                <td>
                  <router-link :to="{name: 'about', hash: '#content', params: {slug: 'integrations'}}">
                    <h5 class="mb-3">Integrations</h5>
                  </router-link>
                </td>
                <td class="d-flex justify-content-end">
                  <i class="mb-3" data-feather="arrow-right"></i>
                </td>
              </tr>

              <tr>
                <td>
                  <router-link :to="{name: 'about', hash: '#content', params: {slug: 'local'}}">
                    <h5 class="mb-3">Local</h5>
                  </router-link>
                </td>
                <td class="d-flex justify-content-end">
                  <i class="mb-3" data-feather="arrow-right"></i>
                </td>
              </tr>
            </table>
          </div>

        </div>
      </div>

    </div>
    
    <hr class="bottom-hr opacity-100">
    
    <div id="content"></div>
    <router-view> </router-view>

    <!-- Content End -->

    <div class="row mt-auto">
      <div class="col-12 mt-lg-5">
        <div class="contacts row">
          
            <div class="col-md-4 offset-md-4 col-sm-6 offset-sm-2">

              <div class="row">
                <div class="col-4 d-flex justify-content-center align-items-center">
                  <a href="https://github.com/AstroSynapseAI" target="_blank" class="btn d-flex justify-content-center align-items-center"><i class="d-block" data-feather="github"></i></a>
                </div>
                <div class="col-4 d-flex justify-content-center align-items-center">
                  <a href="https://discord.gg/yDb4AbEprY" target="_blank" class="btn">
                    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 127.14 96.36">
                      <path class="disc-icon" fill="#fff" d="M107.7,8.07A105.15,105.15,0,0,0,81.47,0a72.06,72.06,0,0,0-3.36,6.83A97.68,97.68,0,0,0,49,6.83,72.37,72.37,0,0,0,45.64,0,105.89,105.89,0,0,0,19.39,8.09C2.79,32.65-1.71,56.6.54,80.21h0A105.73,105.73,0,0,0,32.71,96.36,77.7,77.7,0,0,0,39.6,85.25a68.42,68.42,0,0,1-10.85-5.18c.91-.66,1.8-1.34,2.66-2a75.57,75.57,0,0,0,64.32,0c.87.71,1.76,1.39,2.66,2a68.68,68.68,0,0,1-10.87,5.19,77,77,0,0,0,6.89,11.1A105.25,105.25,0,0,0,126.6,80.22h0C129.24,52.84,122.09,29.11,107.7,8.07ZM42.45,65.69C36.18,65.69,31,60,31,53s5-12.74,11.43-12.74S54,46,53.89,53,48.84,65.69,42.45,65.69Zm42.24,0C78.41,65.69,73.25,60,73.25,53s5-12.74,11.44-12.74S96.23,46,96.12,53,91.08,65.69,84.69,65.69Z"/>
                    </svg>
                    <!-- <img src="../assets/discord-mark-white.svg" alt="discord" height="24"> -->
                  </a>
                </div>
                <div class="col-4 d-flex justify-content-center align-items-center">
                  <a href="https://twitter.com/astrosynapselab" target="_blank" class="btn d-flex justify-content-center align-items-center"><i class="d-block" data-feather="twitter"></i> </a>
                </div>
              </div>

            </div>
          
        </div>

        <div class="row">
          <div class="col-12">
            <p class="text-center"><a href="https://astrosynapse.ai" target="_blank">astrosynapse.ai © 2023</a></p>
          </div>
        </div>
      </div>
    </div>
    
  </div>
  
</template>

<style scoped>
a {
  color: white;
  text-decoration: underline;
}

a:hover {
  color: white;
  text-decoration: underline;
}

nav {
  margin-top: 50px;
  margin-bottom: 6em;
}

nav .btn {
  background-color: black;
  border-color: white;
  border-radius: 0;
  text-decoration: none;
}

nav .btn:hover {
  color: black !important;
  background-color: white;
}

nav .btn:hover .asai-icon line {
  stroke: black !important;
  stroke-width: 10;
}

.logo {
  width: 90%;
}

.sub-menu h5 {
  font-size: 20px;
  font-weight: 400;
  color: white;
}

table i {
  margin-left: 5em;
}

.header {
  margin-top: 150px;
  margin-bottom: 30px;
}

.top-hr {
  margin-bottom: 6em;
}

.bottom-hr {
  margin-top: 6em;
  margin-bottom: 6em;
}

.contacts {
  margin-bottom: 15px;
}

.contacts .btn {
  background-color: black;
  color: white;
  border: 1px solid white;
  width: 60px;
  border-radius: 0;
}

.contacts .btn:hover {
  color: black;
  background-color: white;
}

.btn:hover .disc-icon {
  fill: black;
}

.btn .feather {
  height: 24px;
  width: 24px;
}

@media screen and (max-width: 1024px) {
  .logo {
    width: 100%;
    margin-bottom: 3em;
  }
}

@media only screen and (max-width: 576px) {

  nav {
    margin-top: 10px;
    margin-bottom: 2em;
  }
  .header {
    margin-top: 20px;
  }

  .navbar-menu {
    margin-top: 2em;
  }

  .navbar-menu .row > * {
    padding-right: calc(var(--bs-gutter-x) * .40) !important;
    padding-left: calc(var(--bs-gutter-x) * .40) !important;
  }

  .sub-menu table {
    width: 100%;
  }

  .top-hr {
    margin-bottom: 3em;
  }

  .bottom-hr {
    margin-top: 3em;
    margin-bottom: 3em;
  }


}

</style>