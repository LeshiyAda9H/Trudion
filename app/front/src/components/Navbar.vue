<template>
    <nav class="navbar">
        <div class="navbar-content">
          <RouterLink to="/home" class="nav-home">Home</RouterLink>
          <RouterLink to="/profile" class="nav-link">Profile</RouterLink>
        </div>
    </nav>

</template>

<script  lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "NavbarHeader",
  setup() {

  },
});

</script>

<style>

.navbar {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  display: flex;
  justify-content: space-between;
  padding: 10px 20px;
  background-color: var(--primary-color);
  color: white;
  z-index: 1000;
}

.navbar .nav-link {
  color: white;
  margin: 0 10px;
  transition: color 0.3s;
  
}

.navbar .nav-link:hover {
  color: var(--secondary-color);

}


</style>
