import BaseTemplate from '@/components/parts/base_template/BaseTemplate'

import Index from '@/components/pages/article/Index'
import Create from '@/components/pages/article/create/Create'

export default {
  path: 'article',
  component: BaseTemplate,
  children: [
    {
      path: '',
      component: Index
    },
    {
      path: 'create',
      component: Create
    }
  ]
}