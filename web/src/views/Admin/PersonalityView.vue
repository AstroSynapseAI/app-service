<script setup>
import { Chart, RadarController, CategoryScale, PointElement, LineElement, Title, Tooltip, Legend, RadialLinearScale  } from 'chart.js';
import { ref, onMounted } from 'vue';
import { Form, Field } from 'vee-validate';

// Register the controllers, scales, and elements
Chart.register(RadarController, CategoryScale, PointElement, LineElement, Title, Tooltip, Legend, RadialLinearScale);

let chart = null;
let showModal = ref(false);

const toggleModal = () => {
  showModal.value = !showModal.value;
}

const createChart = () => {
  const ctx = document.getElementById('chartjs-radar').getContext('2d');
  
  if (chart) chart.destroy();  // This will ensure any pre-existing charts are removed before creating a new one

  chart = new Chart(ctx, {
    type: 'radar',
    data: {
      labels: ['Nerouticism', 'Agreeableness', 'Extraversion', 'Conscientiousness', 'Openness to Experience'],
      datasets: [{
        label: 'OCEAN',
        backgroundColor: "rgba(0, 123, 255, 1.0)",
        borderColor: window.theme.primary,
        pointBackgroundColor: window.theme.primary,
        pointBorderColor: "#fff",
        pointHoverBackgroundColor: "#fff",
        pointHoverBorderColor: window.theme.primary,
        data: [46, 73, 44, 83, 93]
      }]
    },
    options: {
      responsive: true, 
      maintainAspectRatio: true,
      elements: {
      line: {
        borderWidth: 3
      }
    },
    scales: {
      r: {
        min: 1,    // minimum value will be 1
        max: 100,  // maximum value will be 100
        ticks: {
          stepSize: 10  // this will add a tick every 10 units
        }
      }
    }
  }
  });
};

onMounted(createChart);

onMounted(() => {
  feather.replace();
});

</script>


<template>
  
          
  <div class="container-fluid p-0">
    
    <!-- Edit Persona Modal -->
    <div class="modal" :class="{ 'd-block': showModal, 'show': showModal }" tabindex="-1">
      <div class="modal-dialog modal-dialog-centered modal-lg">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title">Edit Persona</h5>
            <button type="button" class="btn-close" @click="showModal = false"></button>
          </div>
          <div class="modal-body container">
            <div class="row">
                <div class="col-12">
                  <Field 
                    name="agent_primer" 
                    type="text" 
                    as="textarea" 
                    class="form-control" 
                    rows="8" 
                    placeholder="Describe your persona..."
                  ></Field>
                </div>
              </div>
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-secondary" @click="showModal = false">Close</button>
          </div>
        </div>
      </div>
    </div>

    <h1 class="h3 mb-3">Personality</h1>

    <div class="row">
      <div class="col-12">
        
        <div class="card">

          <div class="card-body">
            <div class="container p-4">
              
              <div class="row mb-5">
                
                <div class="col-6">
                  <div class="form-floating mb-3">
                    <input type="text" class="form-control" id="floatingInput" placeholder="Name your Avatar...">
                    <label for="floatingInput">Avatar name</label>
                  </div>
                </div>

                <div class="col-6">
                  <select class="form-select model-select" aria-label="Select Model">
                    <option selected>GPT-4</option>
                    <option value="3">Claude LLM</option>
                    <option value="1">LLama2</option>
                    <option value="2">Falcon LLM</option>
                  </select>
                </div>

              </div>

              <div class="row mb-5">
                <div class="col-12">
                  <h3>Primer</h3>
                  <Field 
                    name="agent_primer" 
                    type="text" 
                    as="textarea" 
                    class="form-control" 
                    rows="8" 
                    placeholder=""
                  ></Field>
                </div>
              </div>

              <div class="row mb-5">
                <div class="col-12">
                  <h3 class="mb-3">Persona</h3>
                  <table class="table table-bordered table-striped table-hover">
                    <thead>
                      <tr>
                        <th scope="col" class="w-75">Personas</th>
                        <th scope="col" class="w-25">Controls</th>
                      </tr>
                    </thead>
                    <tbody>
                      <tr>
                        <td> Code developer</td>
                        <td>
                          <span @click="toggleModal" style="cursor: pointer; margin-right: 5px;">
                            <i data-feather="edit"></i>
                          </span>
                          <span style="cursor: pointer; margin-right: 5px;">
                            <i data-feather="trash-2"></i>
                          </span>
                        </td>
                        
                      </tr>
                    </tbody>

                  </table>
                </div>
              </div>

              <div class="row">
                <h3>Temperamental Properties</h3>
                <hr>
                <div class="col-12">
                  <h4>Neuroticism</h4>
                </div>
              </div>

              <div class="row">
                <div class="col-6">
                  <label for="customRange1" class="form-label">Whitdrawl</label>
                  <input type="range" class="form-range" id="customRange1">
                </div>
                <div class="col-6">
                  <label for="customRange1" class="form-label">Volatility</label>
                  <input type="range" class="form-range" id="customRange1">
                </div>
              </div>

              <hr>

              <div class="row">
                <div class="col-12">
                  <h4>Agreeableness</h4>
                </div>
              </div>

              <div class="row">
                <div class="col-6">
                  <label for="customRange1" class="form-label">Compassion</label>
                  <input type="range" class="form-range" id="customRange1">
                </div>
                <div class="col-6">
                  <label for="customRange1" class="form-label">Politeness</label>
                  <input type="range" class="form-range" id="customRange1">
                </div>
              </div>

              <hr>

              <div class="row">
                <div class="col-12">
                  <h4>Extraversion</h4>
                </div>
              </div>

              <div class="row">
                <div class="col-6">
                  <label for="customRange1" class="form-label">Enthusiasm</label>
                  <input type="range" class="form-range" id="customRange1">
                </div>
                <div class="col-6">
                  <label for="customRange1" class="form-label">Assertiveness</label>
                  <input type="range" class="form-range" id="customRange1">
                </div>
              </div>

              <hr>

              <div class="row">
                <div class="col-12">
                  <h4>Conscientiousness</h4>
                </div>
              </div>

              <div class="row">
                <div class="col-6">
                  <label for="customRange1" class="form-label">Industriousness</label>
                  <input type="range" class="form-range" id="customRange1">
                </div>
                <div class="col-6">
                  <label for="customRange1" class="form-label">Orderliness</label>
                  <input type="range" class="form-range" id="customRange1">
                </div>
              </div>

              <hr>

              <div class="row">
                <div class="col-12">
                  <h4>Openness to Experience</h4>
                </div>
              </div>

              <div class="row">
                <div class="col-6">
                  <label for="customRange1" class="form-label">Openness</label>
                  <input type="range" class="form-range" id="customRange1">
                </div>
                <div class="col-6">
                  <label for="customRange1" class="form-label">Intellect</label>
                  <input type="range" class="form-range" id="customRange1">
                </div>
                
              </div>

              <hr>

              <div class="row">
                <div class="col-12 d-flex justify-content-center align-items-center">
                  
                  <div style="height: 800px; width: 800px">
                    <canvas id="chartjs-radar"></canvas>
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

  <style scoped>

  .container-fluid {
    color: white;
  }
  .model-select {
    height: 58px;
  }
  </style>