'use client'

import React, { useEffect } from 'react'
import apiClient from '@/apiClient'

function Page() {
  useEffect(() => {
    const fetchData = async () => {
      const res = await apiClient.api.v1.health.$get()
      console.log(res)
    }
    fetchData()
  }, [])
  return <div>This is Health page</div>
}

export default Page
