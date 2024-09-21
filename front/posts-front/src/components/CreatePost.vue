<template>
    <div class="create-post-form">
      <h2>Create a New Post</h2>
      <form @submit.prevent="handleCreatePost">
        <div class="form-group">
          <label for="title">Title</label>
          <input type="text" id="title" v-model="title" required />
        </div>
        <div class="form-group">
          <label for="content">Content</label>
          <textarea id="content" v-model="content" required></textarea>
        </div>
        <button type="submit">Create Post</button>
      </form>
    </div>
  </template>
  
  <script>
  import axiosInstance from '../services/axiosInstance'; 
  
  export default {
    data() {
      return {
        title: '',
        content: '',
      };
    },
    methods: {
      async handleCreatePost() {
        const postData = {
          title: this.title,
          content: this.content,
        };
  
        try {
          const response = await axiosInstance.post('/api/posts', postData); 
          console.log('Post created successfully:', response.data);
          // Reset form fields after successful submission
          this.title = '';
          this.content = '';
          this.$router.push('/profile');
        } catch (error) {
          console.error('Error creating post:', error.message);
          // Handle any error 
        }
      },
    },
  };
  </script>
  
  <style scoped>
  .create-post-form {
    max-width: 800px;
    margin: auto;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
  }

  .form-group {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    padding: 10px;
  }

  input{
  width: 100%;
  padding: 12px;
  border: 1px solid #f7f57c;
  background-color: #242424;
  border-radius: 4px;
  font-size: 16px;
  box-sizing: border-box;
}

  textarea {
    width: 100%;
    height: 150px;
    border-radius: 8px;
    border: 1px solid #f7f57c;
    background-color: #242424;
    padding: 12px;
  }

  button {
    margin-top: 12px;
  }
  </style>
  