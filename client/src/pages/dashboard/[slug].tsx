import DashboardDisplay from '@/components/DashboardDisplay/DashboardDisplay'
import Layout from '@/layouts/Layout'
import { useRouter } from 'next/router'

export default function Dashboard() {
  const router = useRouter()
  const projectId = router.query.slug as string
  return (
    <Layout title='Dashboard'>
      <DashboardDisplay projectId={projectId}/>
    </Layout>
  )
}
