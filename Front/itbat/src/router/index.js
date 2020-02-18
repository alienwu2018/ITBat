import Vue from 'vue'
import Router from 'vue-router'
import index from '@/components/index'
import search from '@/components/search'
import bigcategory from '@/components/BigCategory'
import smallCategory    from '@/components/SmallCategory'
import book from '@/components/book'
import yunpan from '@/components/YunPan'

Vue.use(Router)

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'index',
      component: index
    },
    {
      path: '/search',
      name: 'search',
      component:search
    },
    {
      path: '/category/:BigCategory',
      name: 'bigcategory',
      component :bigcategory,
    },
    {
      path: '/category/:BigCategory/:SmallCategory',
      name: 'smallcategory',
      component :smallCategory,
    },
    {
      path:'/book/:bid',
      name: 'book',
      component:book
    },
    {
      path:'/book/:bid/download/:bid',
      name:'download',
      component:yunpan
    }
  ]
})
