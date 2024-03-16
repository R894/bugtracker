import { useEffect, useState } from 'react'
import BugList from './BugList'
import ProjectsMenu from './ProjectsMenu'
import { useUserContext } from '@/context/UserContext'
import { getBugsByProjectId } from '@/api/bugService'
import { Bug } from '@/api/models/bug'
import CreateProjectModal from './CreateProjectModal'
import { Stack } from '@mui/material'
import CreateBugModal from './CreateBugModal'

const DashboardDisplay = () => {
  const {
    selectedProject,
    setSelectedProject,
    updateUserProjects,
    user,
    token,
    logoutUser,
  } = useUserContext()
  const [bugs, setBugs] = useState<Bug[]>([])

  useEffect(() => {
    updateUserProjects()
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [])

  useEffect(() => {
    if (!token || selectedProject == null) return
    const getBugs = async () => {
      const response = await getBugsByProjectId(selectedProject.id, token)
      if (response && response.code == 401) {
        logoutUser()
        return
      }
      setBugs(response)
    }
    getBugs()
  }, [selectedProject, token, logoutUser])

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
