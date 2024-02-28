import Button from '@mui/material/Button'
import Menu from '@mui/material/Menu'
import MenuItem from '@mui/material/MenuItem'
import { useState, type MouseEvent, type Dispatch, type SetStateAction } from 'react'
import { Project } from '@/api/models/project'

const ProjectsMenu = ({
  projects,
  stateChanger,
  buttonText
}: {
  projects: Project[]
  stateChanger: Dispatch<SetStateAction<Project | null>>
  buttonText: string
}) => {
  const [anchorEl, setAnchorEl] = useState<null | HTMLElement>(null)
  const open = Boolean(anchorEl)

  const handleClick = (event: MouseEvent<HTMLButtonElement>) => {
    setAnchorEl(event.currentTarget)
  }

  const handleClose = (project: Project) => {
    setAnchorEl(null)
    stateChanger(project)
  }
  return (
    <div>
      <Button
        id="basic-button"
        aria-controls={open ? 'basic-menu' : undefined}
        aria-haspopup="true"
        aria-expanded={open ? 'true' : undefined}
        onClick={handleClick}
      >
        {buttonText}
      </Button>
      <Menu
        id="basic-menu"
        anchorEl={anchorEl}
        open={open}
        onClose={() => setAnchorEl(null)}
        MenuListProps={{
          'aria-labelledby': 'basic-button',
        }}
      >
        {projects && projects.map((project) => (
          <MenuItem key={project.id} onClick={() => handleClose(project)}>
            {project.name}
          </MenuItem>
        ))}
      </Menu>
    </div>
  )
}
export default ProjectsMenu