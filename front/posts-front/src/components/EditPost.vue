<template>
    <div class="edit-post-container">
      <h2>Edit Post</h2>
      <form @submit.prevent="updatePost">
        <div class="form-group">
          <label for="title">Title:</label>
          <input type="text" id="title" v-model="post.title" required />
        </div>
  
        <div class="form-group">
          <label for="content">Content:</label>
          <textarea id="content" v-model="post.content" required></textarea>
        </div>
  
        <button type="submit" class="btn">Update Post</button>
        <button type="button" class="btn cancel-btn" @click="goBack">Cancel</button>
      </form>
    </div>
  </template>
  
  <script>
  import axiosInstance from '../services/axiosInstance'; 
  
  export default {
    data() {
      return {
        post: {
          title: '',
          content: '',
        },
        postId: null,
      };
    },
    async created() {
      // Fetch the post details 
      this.postId = this.$route.params.id;
      await this.fetchPost();
    },
    methods: {
      async fetchPost() {
        try {
          const response = await axiosInstance.get(`/api/posts/${this.postId}`);
          this.post = response.data; 
        } catch (error) {
          console.error('Error fetching post:', error.message);
        }
      },
      async updatePost() {
        try {
          const response = await axiosInstance.put(`/api/posts/${this.postId}`, {
            title: this.post.title,
            content: this.post.content,
          });
          console.log('Post updated successfully:', response.data);
          this.$router.push('/profile');
        } catch (error) {
          console.error('Error updating post:', error.message);
        }
      },
      goBack() {
        this.$router.push('/profile');
      },
    },
  };
  </script>
  
  <style scoped>
  .edit-post-container {
    max-width: 600px;
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
    margin-top: 8px;
  }
  
  button.btn {
    background-color: #28a745;
    color: white;
    padding: 10px 20px;
    border: none;
    border-radius: 5px;
    cursor: pointer;
  }
  
  button.btn:hover {
    background-color: #218838;
  }
  
  button.cancel-btn {
    background-color: #dc3545;
    margin-left: 10px;
  }
  
  button.cancel-btn:hover {
    background-color: #c82333;
  }
  </style>
  