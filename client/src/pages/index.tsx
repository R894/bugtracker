import Hero from '@/components/Hero/Hero'
import FeaturesDisplay from '@/components/FeaturesDisplay/FeaturesDisplay'
import Layout from '@/layouts/Layout'

export default function Home() {
  return (
    <Layout>
      <Hero />
      <FeaturesDisplay />
    </Layout>
  )
}
