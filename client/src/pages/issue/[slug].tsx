import BugForm from '@/components/BugForm/BugForm'
import CommentForm from '@/components/Comments/CommentForm'
import CommentList from '@/components/Comments/CommentList'
import { useUserContext } from '@/context/UserContext'
import Layout from '@/layouts/Layout'
import { Stack } from '@mui/material'
import { useRouter } from 'next/router'

export default function Home() {
  const router = useRouter()
  const { token } = useUserContext()
  return (
    <Layout>
      <Stack width={'50%'} minWidth={'320px'} gap={1}>
        {router.query.slug && <BugForm bugId={router.query.slug as string} />}
        <CommentList bugId={router.query.slug as string} token={token} />
        <CommentForm bugId={router.query.slug as string} token={token} />
      </Stack>
    </Layout>
  )
}
