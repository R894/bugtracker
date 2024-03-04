import { getCommentsByBugId } from '@/api/commentService'
import { Comment } from '@/api/models/comment'
import { Grid, Paper } from '@mui/material'
import moment from 'moment'
import { useEffect, useState } from 'react'

const CommentContainer = ({ comment }: { comment: Comment }) => {
  return (
    <Paper sx={{ padding: '40px 20px' }}>
      <Grid container wrap="nowrap" spacing={2}>
        <Grid justifyContent="left" item xs zeroMinWidth>
          <h4 style={{ margin: 0, textAlign: 'left' }}>Michel Michel</h4>
          <p style={{ textAlign: 'left' }}>{comment.content}</p>
          <p style={{ textAlign: 'left' }}>
            {moment(comment.createdAt).format('MMM Do YY')}
          </p>
        </Grid>
      </Grid>
    </Paper>
  )
}

const CommentList = ({ bugId, token }: { bugId: string, token: string }) => {
  const [comments, setComments] = useState<Comment[]>([])
  useEffect(() => {
    if (!token || !bugId) return
    const getComments = async () => {
      console.log(`Using Bug ID: ${bugId} and Token: ${token}`)
      const response = await getCommentsByBugId(bugId, token)
      if (response && response.error != null) {
        return
      }
      console.log(response)
      setComments(response)
    }
    getComments()
  }, [bugId, token])
  return (
    <>
      {comments && comments.map((comment) => (
        <CommentContainer key={comment.id} comment={comment} />
      ))}
    </>
  )
}

export default CommentList
