<template>
    <div class="create-post-form">
      <h2>Create a New Post</h2>
      <form @submit.prevent="handleCreatePost">
        <div>
          <label for="title">Title</label>
          <input type="text" id="title" v-model="title" required />
        </div>
        <div>
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
        } catch (error) {
          console.error('Error creating post:', error.message);
          // Handle any error (e.g., show a notification)
        }
      },
    },
  };
  </script>
  
  <style scoped>
  .create-post-form {
    max-width: 600px;
    margin: auto;
  }
  textarea {
    width: 100%;
    height: 150px;
  }
  </style>
  