import request from '@/utils/request'

export function getThemeList(query) {
  return request({
    url: '/api/v1/themes',
    method: 'get',
    params: query,
  })
}

export function getClientDetail(id) {
  return request({
    url: `/api/v1/clients/${id}/`,
    method: 'get',
  })
}

export function updateClient(data) {
  if (data.group_id === '' || data.group_id === null) {
    data.group_id = 0
  }
  return request({
    url: `/api/v1/clients/${data.id}/`,
    method: 'put',
    data: data,
  })
}

export function deleteClient(id) {
  return request({
    url: `/api/v1/clients/${id}/`,
    method: 'delete',
  })
}

export function batchDeleteClients(params) {
  return request({
    url: `/api/v1/clients/batch_destroy/`,
    method: 'delete',
    params: params,
  })
}

export function getClientFile(uid, query) {
  return request({
    url: `/api/v1/clients/${uid}/file/`,
    method: 'get',
    params: query,
  })
}

export function testFileUpload(uid, path) {
  return request({
    url: `/api/v1/clients/${uid}/file/upload_test/`,
    method: 'get',
    params: {
      uid: uid,
      path: path,
    },
  })
}

export function downloadFile(uid, filePath) {
  return request({
    url: `/api/v1/clients/${uid}/file/download/`,
    method: 'get',
    params: {
      'uid': uid,
      'path': filePath,
    },
    responseType: 'blob',
  })
}

export function renameFile(uid, filePath, newName) {
  return request({
    url: `/api/v1/clients/${uid}/file/rename/`,
    method: 'post',
    data: {
      'uid': uid,
      'path': filePath,
      'new_name': newName,
    },
  })
}

export function moveFile(uid, src, dst) {
  return request({
    url: `/api/v1/clients/${uid}/file/move/`,
    method: 'post',
    data: {
      'uid': uid,
      'src_path': src,
      'dst_folder_path': dst,
    },
  })
}

export function deleteFile(uid, path) {
  return request({
    url: `/api/v1/clients/${uid}/file/delete/`,
    method: 'post',
    data: {
      'uid': uid,
      'path': path,
    },
  })
}

export function mkdirFolder(uid, path) {
  return request({
    url: `/api/v1/clients/${uid}/file/mkdir/`,
    method: 'post',
    data: {
      'uid': uid,
      'path': path,
    },
  })
}

export function getClientGroup(params) {
  return request({
    url: '/api/v1/clientgroups/',
    method: 'get',
    params: params,
  })
}

export function getCountry(params) {
  return request({
    url: '/api/v1/geoip/countries/',
    method: 'get',
    params: params,
  })
}

export function getLoaderVersion() {
  return request({
    url: '/api/v1/clients/loader_version/',
    method: 'get',
  })
}

export function getCPUType() {
  return request({
    url: '/api/v1/clients/cpu_type/',
    method: 'get',
  })
}

export function createClientGroup(data) {
  return request({
    url: `/api/v1/clientgroups/`,
    method: 'post',
    data: data,
  })
}

export function updateClientGroup(data) {
  return request({
    url: `/api/v1/clientgroups/${data.id}/`,
    method: 'put',
    data: data,
  })
}

export function deleteClientGroup(id) {
  return request({
    url: `/api/v1/clientgroups/${id}/`,
    method: 'delete',
  })
}

export function getClientStreamToken(uid) {
  return request({
    url: `/api/v1/clients/${uid}/shell/`,
    method: 'post',
  })
}

export function takeClientOffline(uid) {
  return request({
    url: `/api/v1/clients/${uid}/quit/`,
    method: 'post',
  })
}

export function openProxy(data) {
  return request({
    url: '/api/v1/clients/open_proxy/',
    method: 'post',
    data: data,
  })
}

export function closeProxy(data) {
  return request({
    url: '/api/v1/clients/close_proxy/',
    method: 'post',
    data: data,
  })
}

export function statusProxy(data) {
  return request({
    url: '/api/v1/clients/status_proxy/',
    method: 'post',
    data: data,
  })
}

export function openFRPProxy(data) {
  return request({
    url: '/api/v1/clients/open_frp_proxy/',
    method: 'post',
    data: data,
  })
}

export function closeFRPProxy(data) {
  return request({
    url: '/api/v1/clients/close_frp_proxy/',
    method: 'post',
    data: data,
  })
}

export function statusFRPProxy(data) {
  return request({
    url: '/api/v1/clients/status_frp_proxy/',
    method: 'post',
    data: data,
  })
}

export function openReverseProxy(data) {
  return request({
    url: '/api/v1/clients/open_reverse_proxy/',
    method: 'post',
    data: data,
  })
}

export function closeReverseProxy(data) {
  return request({
    url: '/api/v1/clients/close_reverse_proxy/',
    method: 'post',
    data: data,
  })
}

export function resetReverseProxy(data) {
  return request({
    url: '/api/v1/clients/reset_reverse_proxy/',
    method: 'post',
    data: data,
  })
}

export function resyncClientsReq() {
  return request({
    url: '/api/v1/management/init_client_all/',
    method: 'post',
  })
}
