import Draft from '@/components/pages/draft/Draft'

import Article from '@/components/pages/article/Article'
import Create from '@/components/pages/article/create/Create'

export default {
  path: '/drafts',
  component: Draft,
  children: [
    {
      path: '',
      component: Article
    },
    {
      path: 'new',
      component: Create
    },
    // {
    //   path: ':id',
    //   component: Update
    // }
  ]
}