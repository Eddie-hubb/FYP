import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

/* Layout */
import Layout from '@/layout'

/**
 * Note: sub-menu only appear when route children.length >= 1
 * Detail see: https://panjiachen.github.io/vue-element-admin-site/guide/essentials/router-and-nav.html
 *
 * hidden: true                   if set true, item will not show in the sidebar(default is false)
 * alwaysShow: true               if set true, will always show the root menu
 *                                if not set alwaysShow, when item has more than one children route,
 *                                it will becomes nested mode, otherwise not show the root menu
 * redirect: noRedirect           if set noRedirect will no redirect in the breadcrumb
 * name:'router-name'             the name is used by <keep-alive> (must set!!!)
 * meta : {
    roles: ['admin','editor']    control the page roles (you can set multiple roles)
    title: 'title'               the name show in sidebar and breadcrumb (recommend set)
    icon: 'svg-name'             the icon show in the sidebar
    breadcrumb: false            if set false, the item will hidden in breadcrumb(default is true)
    activeMenu: '/example/list'  if set path, the sidebar will highlight the path you set
  }
 */

/**
 * constantRoutes
 * a base page that does not have permission requirements
 * all roles can be accessed
 */
export const constantRoutes = [{
  path: '/login',
  component: () => import('@/views/login/index'),
  hidden: true
},

{
  path: '/404',
  component: () => import('@/views/404'),
  hidden: true
},

{
  path: '/',
  component: Layout,
  redirect: '/machine-learning',
  children: [{
    path: 'machine-learning',
    name: 'Machine Learning',
    component: () => import('@/views/ml/index'),
    meta: {
      title: 'Machine Learning',
      icon: 'realestate'
    }
  }]
}
]

/**
 * asyncRoutes
 * the routes that need to be dynamically loaded based on user roles
 */
export const asyncRoutes = [
  {
    path: '/portfolio',
    component: Layout,
    redirect: '/portfolio/all',
    name: 'Selling',
    alwaysShow: true,
    meta: {
      title: 'Portfolio',
      icon: 'selling'
    },
    children: [{
      path: 'all',
      name: 'SellingAll',
      component: () => import('@/views/selling/all/index'),
      meta: {
        title: 'All Portfolio',
        icon: 'sellingAll'
      }
    },
    {
      path: 'create',
      name: 'SellingBuy',
      component: () => import('@/views/selling/buy/index'),
      meta: {
        title: 'Create Portfolio',
        icon: 'sellingBuy'
      }
    }
    ]
  },
  {
    path: '/transaction',
    component: Layout,
    redirect: '/transaction/information',
    name: 'Donating',
    alwaysShow: true,
    meta: {
      title: 'Transaction',
      icon: 'donating'
    },
    children: [{
      path: 'information',
      name: 'Information',
      component: () => import('@/views/donating/all/index'),
      meta: {
        title: 'Transaction Information',
        icon: 'donatingAll'
      }
    },
    {
      path: 'money',
      name: 'Money',
      component: () => import('@/views/donating/donor/index'),
      meta: {
        title: 'Money',
        icon: 'donatingDonor'
      }
    }, {
      path: 'commodity',
      name: 'Commodity',
      component: () => import('@/views/donating/grantee/index'),
      meta: {
        title: 'Commodity',
        icon: 'donatingGrantee'
      }
    },
    {
      path: 'service',
      name: 'ServiceCharge',
      component: () => import('@/views/donating/service/index'),
      meta: {
        title: 'Service Charge',
        icon: 'donatingGrantee'
      }
    },
    {
      path: 'redemption',
      name: 'RedemptionFee',
      component: () => import('@/views/donating/redemption/index'),
      meta: {
        title: 'Redemption Fee',
        icon: 'donatingGrantee'
      }
    }
    ]
  },
  {
    path: '/addRealestate',
    component: Layout,
    meta: {
      roles: ['admin']
    },
    children: [{
      path: '/addRealestate',
      name: 'AddRealestate',
      component: () => import('@/views/realestate/add/index'),
      meta: {
        title: '????????????',
        icon: 'addRealestate'
      }
    }]
  },
  // 404 page must be placed at the end !!!
  {
    path: '*',
    redirect: '/404',
    hidden: true
  }
]

const createRouter = () => new Router({
  base: '/web',
  // mode: 'history', // require service support
  scrollBehavior: () => ({
    y: 0
  }),
  routes: constantRoutes
})

const router = createRouter()

// Detail see: https://github.com/vuejs/vue-router/issues/1234#issuecomment-357941465
export function resetRouter() {
  const newRouter = createRouter()
  router.matcher = newRouter.matcher // reset router
}

export default router
