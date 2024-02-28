import { useContext, useEffect, useState } from 'react'
import BugList from './BugList'
import ProjectsMenu from './ProjectsMenu'
import { UserContext, UserContextType } from '@/context/UserContext'
import { getBugsByProjectId } from '@/api/bugService'
import { Bug } from '@/api/models/bug'
import CreateProjectModal from './CreateProjectModal'
import { Stack } from '@mui/material'
import { Project } from '@/api/models/project'
import CreateBugModal from './CreateBugModal'

const DashboardDisplay = () => {
  const { updateUserProjects, user, token } = useContext(
    UserContext,
  ) as UserContextType
  const [selectedProject, setSelectedProject] = useState<Project | null>(null)
  const [bugs, setBugs] = useState<Bug[]>([])

  useEffect(() => {
    updateUserProjects()
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [])

  useEffect(() => {
    if (!token || selectedProject == null) return
    const getBugs = async () => {
      const response = await getBugsByProjectId(selectedProject.id, token)
      console.log('project id', selectedProject.id)
      console.log(response)
      setBugs(response)
    }
    getBugs()
  }, [selectedProject, token])

  return (
    <>
      <div style={{ padding: '16px' }}>
        <Stack direction={'row'}>
          <ProjectsMenu
            buttonText={
              selectedProject ? selectedProject.name : 'Select Project'
            }
            projects={user ? user.projects : []}
            stateChanger={setSelectedProject}
          />
          <CreateProjectModal />
          {selectedProject ? (
            <CreateBugModal projectId={selectedProject.id} />
          ) : null}
        </Stack>
        <BugList bugs={bugs} />
      </div>
    </>
  )
}

export default DashboardDisplay
