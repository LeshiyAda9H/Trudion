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
  /* display: flex;
  justify-content: flex-end; */
  padding: 5px;
  background-color: #fff;
}


.navbar-content {
  display: flex;
  align-items: center;
  width: 100%;
}

.nav-home{
  top: 0;
  left: 0;
  padding: 5px 10px;
  text-decoration: none;
  color: #007bff;
  border-radius: 5px;
}
.nav-home:hover {
  background-color: #e2e6ea;
  color: #0056b3;
}

.nav-link {

  padding: 5px 10px;
  text-decoration: none;
  color: #007bff;
  border-radius: 5px;
}

.nav-link:hover {
  background-color: #e2e6ea;
  color: #0056b3;
}

</style>
