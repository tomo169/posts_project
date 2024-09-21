<template>
    <div class="profile">
      <h2>My Posts</h2>
      <div v-if="posts.length > 0">
        <PostCard 
          v-for="post in posts" 
          :key="post.id" :post="post" 
          :showAuthor="false"  
          :showActions="true" 
          @edit-post="editPost" 
          @delete-post="handleDeletePost"
          />
      </div>
      <p v-else>No posts found for this user.</p>
    </div>
  </template>
  
  <script>
  import axiosInstance from '../services/axiosInstance';  
  import { jwtDecode } from "jwt-decode";
  import PostCard from './PostCard.vue'; 
  
  export default {
    components: {
    PostCard, 
  },
    data() {
      return {
        posts: [],
        userID: null,
      };
    },
    async mounted() {
      // Decode JWT to get the userID
      const token = localStorage.getItem('authToken');
      if (token) {
        const decodedToken = jwtDecode(token);
        this.userID = decodedToken.userID;  
  
        // Fetch posts for the logged-in user
        await this.fetchUserPosts();
      }
    },
    methods: {
      async fetchUserPosts() {
        try {
          const response = await axiosInstance.get(`/api/users/${this.userID}/posts`);
          this.posts = response.data;  
          console.log('Fetched posts:', this.posts);
        } catch (error) {
          console.error('Error fetching posts:', error.message);
        }
      },

      editPost(post) {
      // Navigate to the EditPost component with the postId
      this.$router.push({ name: 'edit-post', params: { id: post.id } });
    },

    async handleDeletePost(postId) {
      try {
        await axiosInstance.delete(`/api/posts/${postId}`);
        // Remove the post from the posts array
        this.posts = this.posts.filter((post) => post.id !== postId);
      } catch (error) {
        console.error('Error deleting post:', error.message);
      }
    },
  },
};
  </script>
  
  <style scoped>
  .profile {
    max-width: 800px;
    margin: auto;
  }
  </style>
  