import { useAuthStore } from '@/stores/auth.store';

export const fetchWrapper = {
    get: request('GET'),
    post: request('POST'),
    put: request('PUT'),
    delete: request('DELETE')
};

function request(method) {
  return (url, body) => {
      const requestOptions = {
          method,
          headers: authHeader(url)
      };
      if (body) {
          requestOptions.headers['Content-Type'] = 'application/json';
          requestOptions.body = JSON.stringify(body);
      }
      return fetch(url, requestOptions).then(handleResponse);
  }
}

function authHeader(url) {
  let user = localStorage.getItem('user')
  const isApiUrl = url.startsWith(import.meta.env.VITE_API_URL);
  if (user && (user.apiToken && isApiUrl)) {
      return { Authorization: `Bearer ${user.apiToken}` };
  } else {
      return {};
  }
}

async function handleResponse(response) {
    const isJson = response.headers?.get('content-type')?.includes('application/json');
    const data = isJson ? await response.json() : null;

    // check for error response
    if (!response.ok) {
        // const { user, logout } = useAuthStore();
        // if ([401, 403].includes(response.status) && user) {
        //     // auto logout if 401 Unauthorized or 403 Forbidden response returned from api
        //     logout();
        // }

        // get error message from body or default to response status
        
        const error = (data && data.message) || response.status;
        return Promise.reject(data || error);
    }

    return data;
}