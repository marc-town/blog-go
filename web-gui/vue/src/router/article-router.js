import BaseTemplate from '@/components/parts/base_template/BaseTemplate'

import Article from '@/components/pages/article/Article'
import Create from '@/components/pages/article/create/Create'

export default {
  path: '/article',
  component: BaseTemplate,
  children: [
    {
      path: '',
      component: Article
    },
    {
      path: 'create',
      component: Create
    }
  ]
}