<template>
    <nav class="navbar">
      <ul class="nav-left">
        <li><router-link to="/">Home</router-link></li>
        <li><router-link to="/posts">Posts</router-link></li>
        <li><router-link to="/profile">Profile</router-link></li>
      </ul>
      <ul class="nav-right">
        <li v-if="isLoggedIn"><a @click="handleLogout">Logout</a></li>
        <li v-if="!isLoggedIn"><router-link to="/login">Login</router-link></li>
        <li v-if="!isLoggedIn"><router-link to="/register">Sign Up</router-link></li>
      </ul>
    </nav>
  </template>
  
  <script>
  export default {
    data() {
      return {
        isLoggedIn: false,
      };
    },
    mounted() {
      // Check if the user is logged in by looking for a token in localStorage
      this.isLoggedIn = !!localStorage.getItem('authToken');
    },
    methods: {
      handleLogout() {
        // Remove the JWT token from localStorage
        localStorage.removeItem('authToken');
        this.isLoggedIn = false;
        // Redirect to login page
        this.$router.push('/login');
      },
    },
    watch: {
      // Watch for changes in the route or token
      $route() {
        this.isLoggedIn = !!localStorage.getItem('authToken');
      }
    }
  };
  </script>
  
  <style scoped>
  .navbar {
    display: flex;
    justify-content: space-between;
    padding: 10px;
    background-color: #333;
    color: white;
  }
  
  .nav-left, .nav-right {
    list-style-type: none;
    display: flex;
    gap: 15px;
  }
  
  .nav-left li, .nav-right li {
    cursor: pointer;
  }
  
  .nav-left li a, .nav-right li a {
    color: white;
    text-decoration: none;
  }
  
  .nav-left li a:hover, .nav-right li a:hover {
    text-decoration: underline;
  }
  </style>
  