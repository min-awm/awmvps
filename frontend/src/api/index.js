export default {
  // auth
  LOGIN: "/auth/login",
  REGISTER: "/auth/register",
  USER_INFO: "/auth/user-info",
  CHANGE_USER_INFO: "/auth/change-user-info",
  PACKAGE_MANAGER: "/auth/package-manager",

  // filemanager
  CHECK_PATH: "/filemanager/check-path",
  GET_LIST: "/filemanager/list",
  RENAME_FILE: "/filemanager/rename",

  // database
  STATUS_MARIADB: "/databases/status-mariadb",
  INSTALL_MARIADB: "/databases/install-mariadb",
  START_MARIADB: "/databases/start-mariadb",
  STOP_MARIADB: "/databases/stop-mariadb",
  LIST_USER_MARIADB: "/databases/list-user-mariadb",
  CREATE_USER_MARIADB: "/databases/create-user-mariadb",
  DELETE_USER_MARIADB: "/databases/delete-user-mariadb",
  REMOVE_MARIADB: "/databases/remove-mariadb",

  // docker
  STATUS_DOCKER: "/docker/status",
  LIST_CONTAINER_DOCKER: "/docker/list-container",
  STOP_CONTAINER_DOCKER: "/docker/stop-container",
  START_CONTAINER_DOCKER: "/docker/start-container",
  REMOVE_CONTAINER_DOCKER: "/docker/remove-container",
  
  LIST_IMAGE_DOCKER: "/docker/list-image",
  REMOVE_IMAGE_DOCKER: "/docker/remove-image",

  LIST_VOLUME_DOCKER: "/docker/list-volume",
  REMOVE_VOLUME_DOCKER: "/docker/remove-volume",

  // nginx
  INSTALL_NGINX: "/nginx/install",
  STATUS_NGINX: "/nginx/status",
  START_NGINX: "/nginx/start",
  STOP_NGINX: "/nginx/stop",
  LIST_CONF_NGINX: "/nginx/list-conf",
  FILE_CONF_NGINX: "/nginx/file-conf",
  CREATE_CONF_NGINX: "/nginx/create-conf",
  REMOVE_CONF_NGINX: "/nginx/remove-conf",
  SAVE_CONF_NGINX: "/nginx/save-conf",
  CHECK_CONF_NGINX: "/nginx/check-conf",

  // block ip
  LIST_BLOCK_IP: "/security/list-block-ip",
  ADD_BLOCK_IP: "/security/add-block-ip",
  REMOVE_BLOCK_IP: "/security/remove-block-ip",

  // port
  LIST_PORT: "/security/list-port",
  ADD_PORT: "/security/add-port",
  REMOVE_PORT: "/security/remove-port",

  // ddos
  ADD_ROLE: "/security/add-role"
};
