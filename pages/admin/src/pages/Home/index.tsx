import { Button, Loading, Table } from '@geist-ui/react'
import React, { useEffect, useState } from 'react'
import { RouteChildrenProps } from 'react-router'
import { getProjects } from '../../api'
import { Project } from '../../types'

const HomePage: React.FC<RouteChildrenProps> = ({ match }) => {
  const [projects, setProjects] = useState<Project[]>([])
  const [loading, setLoading] = useState(false)

  useEffect(() => {
    setLoading(true)
    getProjects().then(({ error, data }) => {
      setProjects(
        error
          ? []
          : data!.map(item => ({
              ...item,
              operation: (actions, rowData) => {
                return (
                  <Button type="error" auto size="mini">
                    删除
                  </Button>
                )
              }
            }))
      )
    })
  }, [])

  return (
    <>
      {loading && <Loading />}
      <Table data={projects}>
        <Table.Column prop="name" label="Name" />
        <Table.Column prop="operation" label="Operation" width={150} />
      </Table>
    </>
  )
}

export default HomePage
