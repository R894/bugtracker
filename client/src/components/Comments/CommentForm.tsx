import { createComment } from '@/api/commentService'
import { Button, TextField } from '@mui/material'
import { useState } from 'react'

const CommentForm = ({ bugId, token }: { bugId: string; token: string }) => {
  const [commentText, setCommentText] = useState('')
  const handleSubmitComment = async () => {
    if (!token || !bugId) return
    const response = await createComment({ bugId, content: commentText }, token)
    console.log(response)
  }
  return (
    <>
      <TextField
        multiline
        fullWidth
        minRows={2}
        id="outlined-multilined"
        placeholder="Write comment here..."
        value={commentText}
        onChange={(e) => {
          setCommentText(e.target.value)
        }}
      />
      <Button onClick={handleSubmitComment}>Submit</Button>
    </>
  )
}

export default CommentForm
