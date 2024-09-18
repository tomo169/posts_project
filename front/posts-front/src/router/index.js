import {createRouter, createWebHistory} from 'vue-router';
import Login from '../components/Login.vue';
import Register from '../components/Register.vue';
import CreatePost from '../components/CreatePost.vue';

function isAuthenticated() {
    return !!localStorage.getItem('authToken');
  }

const routes = [
    { path: '/login', component: Login },
    { path: '/Register', component: Register },

    { 
        path: '/create-post', 
        component: CreatePost,
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
    next();
  });

export default router;