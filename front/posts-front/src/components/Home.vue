<template>
    <div class="home-container">
      <h1>Welcome to Demo Posts App!</h1>
      <p class="description">
        This is a platform where you can share your thoughts and write posts. 
        To get started, please log in or create an account.
      </p>
      <div class="action-buttons">
        <button @click="goToLogin" class="btn login-btn">Login</button>
        <router-link to="/register" class="register-link">Don't have an account? Register here</router-link>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    methods: {
      goToLogin() {
        this.$router.push('/login');
      },
    },
  };
  </script>
  
  <style scoped>
  .home-container {
    max-width: 600px;
    margin: 100px auto;
    text-align: center;
  }
  
  h1 {
    font-size: 36px;
    color: #17beb6;
    margin-bottom: 20px;
  }
  
  .description {
    font-size: 18px;
    color: #dad3d3;
    margin-bottom: 40px;
  }
  
  .action-buttons {
    display: flex;
    flex-direction: column;
    align-items: center;
  }

  button {
    margin-bottom: 20px;
  }

  
  .register-link {
    color: #fdff76;
    text-decoration: none;
    font-size: 16px;
  }
  
  .register-link:hover {
    text-decoration: underline;
  }
  </style>
  