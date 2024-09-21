<template>
  <div class="post-card">
    <div class="post-header">
      <span v-if="showAuthor" class="username">{{ post.authorName || 'Unknown Author' }}</span> 
    </div>
    <h3>{{ post.title }}</h3>
    <p>{{ post.content }}</p>
    <div v-if="showActions" class="post-actions">
      <button @click="editPost">Edit</button>
      <button @click="deletePost" class="delete">Delete</button>
    </div>
  </div>
</template>
  
  <script>
  export default {
    props: {
      post: {
        type: Object,
        required: true,
      },
      showAuthor: {
        type: Boolean,
        default: true,
      },
      showActions: {
        type: Boolean,
        default: false, 
      },
    },
    methods: {
    editPost() {
      // Emit an event for edit, passing the post as payload
      this.$emit('edit-post', this.post);
    },
    async deletePost() {
      // Confirm before deleting
      if (confirm("Are you sure you want to delete this post?")) {
        try {
          await this.$emit('delete-post', this.post.id);
        } catch (error) {
          console.error('Error deleting post:', error);
        }
      }
    },
  },
};
  </script>
  
  <style scoped>
  .post-card {
    border: 1px solid #17beb6;
    border-radius: 8px;
    padding: 16px;
    margin-bottom: 16px;
    background-color: #333;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  }
  
  .post-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 8px;
}

.post-actions {
  display: flex;
  justify-content: flex-end;
  margin-top: 20px;
}

button {
  margin-left: 10px;
  padding: 8px 12px;
  color: white;
  cursor: pointer;
}

button.delete {
  background-color: #dc3545;
}

button.delete:hover {
  background-color: #c82333;
}

.username {
  font-weight: bold;
  color: white;
}

h3 {
  margin-top: 0;
  color: white;
}

p {
  color: white;
}
</style>
  