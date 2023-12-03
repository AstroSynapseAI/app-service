
<template>
  <nav class="navbar navbar-expand navbar-light navbar-bg">
    <a class="sidebar-toggle js-sidebar-toggle">
      <i class="hamburger align-self-center"></i>
    </a>

    <div class="navbar-collapse collapse">
      <ul class="navbar-nav navbar-align">
        
        <li class="nav-item dropdown">	
          <a class="nav-icon dropdown-toggle d-inline-block d-sm-none" href="#" data-bs-toggle="dropdown">
            <i class="align-middle" data-feather="settings"></i>
          </a>

          <a class="nav-link dropdown-toggle d-none d-sm-inline-block" href="#" data-bs-toggle="dropdown">
            <span class="text-dark">Account</span>
          </a>
          
          <div class="dropdown-menu dropdown-menu-end">
            <router-link to="/account" class="dropdown-item"><i class="align-middle me-1" data-feather="user"></i> Profile</router-link>
            
            <div class="dropdown-divider"></div>
            <a class="dropdown-item" href="#"><i class="align-middle me-1" data-feather="log-out"></i>Log out</a>
          </div>
        </li>

      </ul>
    
    </div>
  </nav>
</template>

<style scoped>
.navbar, .sidebar-toggle {
  background-color: #221f1f;
  color: white;
  
}

.navbar {
  border-bottom: 4px solid white;
}




</style>