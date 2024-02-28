import Box from '@mui/material/Box'
import Button from '@mui/material/Button'
import Modal from '@mui/material/Modal'
import { TextField, Typography } from '@mui/material'
import { useContext, useState } from 'react'
import { createProject } from '@/api/projectService'
import { UserContext, UserContextType } from '@/context/UserContext'

const style = {
  position: 'absolute' as 'absolute',
  top: '50%',
  left: '50%',
  transform: 'translate(-50%, -50%)',
  width: 400,
  bgcolor: 'background.paper',
  boxShadow: 24,
  display: 'flex',
  flexDirection: 'column',
  p: 4,
}

const CreateProjectModal = () => {
  const [open, setOpen] = useState(false)
  const [projectName, setProjectName] = useState('')
  const [projectDescription, setProjectDescription] = useState('')
  const {token} = useContext(UserContext) as UserContextType
  const handleOpen = () => setOpen(true)
  const handleClose = () => setOpen(false)

  const postProject = async () => {
    await createProject({name: projectName, description: projectDescription}, token)
  }

  return (
    <div>
      <Button onClick={handleOpen}>New Project</Button>
      <Modal
        open={open}
        onClose={handleClose}
        aria-labelledby="modal-modal-title"
        aria-describedby="modal-modal-description"
      >
        <Box sx={style}>
          <Typography variant="h6" sx={{ textAlign: 'center' }}>
            Create New Project
          </Typography>
          <TextField
            variant="standard"
            value={projectName}
            onChange={(e) => setProjectName(e.target.value)}
            label="Name"
          />
          <TextField
            variant="standard"
            value={projectDescription}
            onChange={(e) => setProjectDescription(e.target.value)}
            label="Description"
          />
          <Button onClick={postProject}>Create</Button>
        </Box>
      </Modal>
    </div>
  )
}

export default CreateProjectModal
