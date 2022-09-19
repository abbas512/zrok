/** @module types */
// Auto-generated, edits will be overwritten

/**
 * @typedef accountRequest
 * @memberof module:types
 * 
 * @property {string} email 
 * @property {string} password 
 */

/**
 * @typedef accountResponse
 * @memberof module:types
 * 
 * @property {string} token 
 */

/**
 * @typedef authUser
 * @memberof module:types
 * 
 * @property {string} username 
 * @property {string} password 
 */

/**
 * @typedef enableRequest
 * @memberof module:types
 * 
 * @property {string} description 
 * @property {string} host 
 */

/**
 * @typedef enableResponse
 * @memberof module:types
 * 
 * @property {string} identity 
 * @property {string} cfg 
 */

/**
 * @typedef disableRequest
 * @memberof module:types
 * 
 * @property {string} identity 
 */

/**
 * @typedef environment
 * @memberof module:types
 * 
 * @property {string} description 
 * @property {string} host 
 * @property {string} address 
 * @property {string} zitiIdentityId 
 * @property {boolean} active 
 * @property {string} createdAt 
 * @property {string} updatedAt 
 */

/**
 * @typedef environmentServices
 * @memberof module:types
 * 
 * @property {module:types.environment} environment 
 * @property {module:types.services} services 
 */

/**
 * @typedef loginRequest
 * @memberof module:types
 * 
 * @property {string} email 
 * @property {string} password 
 */

/**
 * @typedef principal
 * @memberof module:types
 * 
 * @property {number} id 
 * @property {string} email 
 * @property {string} token 
 */

/**
 * @typedef service
 * @memberof module:types
 * 
 * @property {string} zitiServiceId 
 * @property {string} endpoint 
 * @property {boolean} active 
 * @property {string} createdAt 
 * @property {string} updatedAt 
 */

/**
 * @typedef tunnelRequest
 * @memberof module:types
 * 
 * @property {string} zitiIdentityId 
 * @property {string} endpoint 
 * @property {string} authScheme 
 * @property {module:types.authUser[]} authUsers 
 */

/**
 * @typedef tunnelResponse
 * @memberof module:types
 * 
 * @property {string} proxyEndpoint 
 * @property {string} service 
 */

/**
 * @typedef untunnelRequest
 * @memberof module:types
 * 
 * @property {string} zitiIdentityId 
 * @property {string} service 
 */

/**
 * @typedef verifyRequest
 * @memberof module:types
 * 
 * @property {string} token 
 */

/**
 * @typedef verifyResponse
 * @memberof module:types
 * 
 * @property {string} email 
 */
