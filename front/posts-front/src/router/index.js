import {createRouter, createWebHistory} from 'vue-router';
import Home from '../components/Home.vue';
import Login from '../components/Login.vue';
import Register from '../components/Register.vue';
import CreatePost from '../components/CreatePost.vue';
import Profile from '../components/Profile.vue';
import Posts from '../components/Posts.vue';
import EditPost from '../components/EditPost.vue';

function isAuthenticated() {
    return !!localStorage.getItem('authToken');
  }

const routes = [
    { path: '/', component: Home},
    { path: '/Posts', component: Posts },
    
    { 
      path: '/login', 
      component: Login, 
      meta: { requiresGuest: true }, 
    },

    { 
      path: '/Register', 
      component: Register,
      meta: { requiresGuest: true },
    },

    { 
        path: '/create-post', 
        component: CreatePost,
        meta: { requiresAuth: true}
    },

    { 
      path: '/edit-post/:id',
      name: 'edit-post', 
      component: EditPost,
      meta: { requiresAuth: true}
  },

    { 
      path: '/profile', 
      component: Profile,
      meta: { requiresAuth: true}
  },
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

// Global navigation guard
router.beforeEach((to, from, next) => {
    if (to.matched.some(record => record.meta.requiresAuth)) {

      if (!isAuthenticated()) {

        return next({ path: '/login' });
      }
    }
    if (to.matched.some(record => record.meta.requiresGuest)) {
      if (isAuthenticated()) {
          return next({ path: '/profile' });
      }
    }
    next();
  });

export default router;