import { configHTTP, configInterceptor } from './config/request';
import { configValidator } from './config/validate'
import { webSocket } from "./helpers/webSocket"

configHTTP();
configInterceptor();
configValidator();
webSocket();
