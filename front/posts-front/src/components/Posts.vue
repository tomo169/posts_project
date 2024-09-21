<template>
    <div class="posts-container" @scroll="handleScroll">
      <h2>Latest Posts</h2>
      <div v-if="posts.length > 0">
        <PostCard v-for="post in posts" :key="post.id" :post="post" />
      </div>
      <p v-else>No posts available.</p>
      <div v-if="loading" class="loading-spinner">
        <p>Loading...</p>
      </div>
    </div>
  </template>
  
  <script>
  import PostCard from './PostCard.vue'; 
  import axiosInstance from '../services/axiosInstance';
  
  export default {
    components: {
      PostCard,
    },
    data() {
      return {
        posts: [],      
        offset: 0,      
        limit: 10,      
        loading: false,  
        hasMorePosts: true, 
      };
    },
    async mounted() {
      // Fetch the initial set of posts when the component mounts
      await this.fetchPosts();
    },
    methods: {
      async fetchPosts() {
        if (this.loading || !this.hasMorePosts) return; // Avoid fetching if already loading or no more posts available
  
        this.loading = true;
  
        try {
          // Fetch posts from the backend using the offset and limit
          const response = await axiosInstance.get(`/posts`, {
            params: {
              limit: this.limit,
              offset: this.offset,
            },
          });
  
          const newPosts = response.data;
  
          // If fewer posts are returned than the limit, it means no more posts are available
          if (newPosts.length < this.limit) {
            this.hasMorePosts = false;
          }
  
          // Add the newly fetched posts to the existing posts array
          this.posts = [...this.posts, ...newPosts];
  
          // Increment the offset for the next batch of posts
          this.offset += this.limit;
        } catch (error) {
          console.error('Error fetching posts:', error.message);
        } finally {
          this.loading = false;
        }
      },
      handleScroll(event) {
        const container = event.target;
  
        // Detect if the user has scrolled to the bottom of the container
        if (container.scrollTop + container.clientHeight >= container.scrollHeight - 10) {
          this.fetchPosts();
        }
      },
    },
  };
  </script>
  
  <style scoped>
.posts-container {
  max-width: 800px;
  margin: auto;
  height: 80vh; 
  overflow-y: scroll; 
  padding: 20px;
  border: 1px solid #dadd17;
  border-radius: 8px;
  scroll-behavior: smooth; 
  margin-top: 30px;
  
  scrollbar-width: none; 
  -ms-overflow-style: none; 
}

.posts-container::-webkit-scrollbar {
  display: none; 
}
  
  .loading-spinner {
    text-align: center;
    padding: 20px;
  }
  
  h2 {
    text-align: center;
    margin-bottom: 20px;
  }
  </style>
  