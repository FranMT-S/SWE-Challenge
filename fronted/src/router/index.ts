import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import HomeView from '../views/Home/HomeView.vue'
import {ERouterName} from "@/helpers/enums/RouterName.enum"

const routes: Array<RouteRecordRaw> = [
  {
    path: '/home',
    name: ERouterName.HOME,
    redirect:{name:ERouterName.MAILS},
    component: HomeView,
    children:[
      {
        path: 'inbox',
        name: ERouterName.INBOX,
        
        children:[
            {
                path: '',
                name: ERouterName.MAILS,
                component: () => import(/* webpackChunkName: "about" */ '../views/Inbox/InboxView.vue'),  
            }
        ]
      },
     
    
    ]
  },
  {
    path: '/',
    redirect: '/home'
  },

  {
    path: '/:pathMatch(.*)',
    redirect: '/home'
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
