import DashboardDisplay from '@/components/DashboardDisplay/DashboardDisplay'
import Layout from '@/layouts/Layout'
import Head from 'next/head'

export default function Dashboard() {
  return (
    <>
      <Head>
        <title>Dashboard</title>
        <link rel="icon" href="/favicon.ico" />
        <meta
          name="viewport"
          content="width=device-width, initial-scale=1, user-scalable=no, shrink-to-fit=no, viewport-fit=cover"
        />
      </Head>
      <main>
        <Layout>
          <DashboardDisplay />
        </Layout>
      </main>
    </>
  )
}
