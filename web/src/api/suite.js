import request from '@/utils/request'

export function getSuiteList(query) {
  return request({
    url: '/api/v1/suites',
    method: 'get',
    params: query,
  })
}

export function getSuiteDetail(suiteId) {
  return request({
    url: `/api/v1/suites/${suiteId}`,
    method: 'get',
  })
}
