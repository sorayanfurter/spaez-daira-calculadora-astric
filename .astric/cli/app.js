import { cp, stat, readFile, writeFile, readdir } from 'fs/promises';

const folderCLI = './.astric/cli/structure/';
const folderFront = '/src/app/';
const folderEndpoint = './BackEnd/api/'
let routesAppFile = `./FrontEnd/src/routes.svelte`;
let routesModFile = `./FrontEnd/src/app`;

async function copy(from, to) {
    try {
        const checkDirFrontEnd = (await readdir(`./FrontEnd${folderFront}`))
        // const checkDirBackEnd = (await readdir(folderEndpoint))

        let newTo = to.split('/')

        if (checkDirFrontEnd.find(d => d === to)) {
            throw new Error("DIRECTORIO DE APP EN FrontEnd EXISTENTE")
        }

        // if (checkDirBackEnd.find(d => d === (newTo[0] ? newTo[0] : to))) {
        //     const checkDirBackEndApi = (await readdir(`${folderEndpoint}/${newTo[0]}/controllers/`))
        //     if (checkDirBackEndApi.find(d => d === `${newTo[1]}.go`)) {
        //         throw new Error("DIRECTORIO DE MOD EN BACKEND EXISTENTE")
        //     }

        //     // let newRoutes = await readFile(folderCLI + `endpoint/routes/routes.go`,"utf-8")

        //     // newRoutes = newRoutes.replace(`func rutas(ruta routes.TipoRuta) {`, `func rutas(ruta routes.TipoRuta) {\n` + `ruta("/${newTo[1]}", controllers.${newTo[1]}, "GET", "Ruta de prueba"),`)

        //     // await writeFile(`${folderEndpoint}/${newTo[0]}/routes/routes.go`, newRoutes)

        //     let newController = await readFile(folderCLI + `endpoint/controllers/prueba.go`, "utf-8")

        //     newController = newController.split('func Prueba(w http.ResponseWriter, r *http.Request) {');

        //     newController.splice(newController[1].indexOf(), 0, `func ${newTo[1]}(w http.ResponseWriter, r *http.Request) {`);

        //     // await cp(folderCLI + `endpoint/controllers/prueba.go`, `${folderEndpoint}/${newTo[0]}/controllers/${newTo[0]}.go`, { recursive: true });
        //     let finalController = newController.join('');
        //     await writeFile(`${folderEndpoint}/${newTo[0]}/controllers/${newTo[1]}.go`, finalController)

        // }


        await cp(folderCLI + from, `./FrontEnd${folderFront}${to}`, { recursive: true });
        console.log('DIRECTORIO DE APP EN FrontEnd:', `./FrontEnd${folderFront}${to}`)

        // await cp(folderCLI + 'endpoint', folderEndpoint + (newTo[0] ? newTo[0] : to), { recursive: true });
        // console.log('DIRECTORIO DE APP EN BACKEND:', folderCLI + 'endpoint', folderEndpoint + (newTo[0] ? newTo[0] : to))

        // let newRoutes = await readFile(folderCLI + `endpoint/routes/routes.go`, "utf-8")

        // newRoutes = newRoutes.replace("ASTRIC/BackEnd/api/CAMBIARNOMBRE/controllers", `ASTRIC/BackEnd/api/${newTo[0]}/controllers`)

        // await writeFile(`${folderEndpoint}/${newTo[0]}/routes/routes.go`, newRoutes)

        let separatorApp = await readFile(routesAppFile, 'utf-8');
        let separatorMod = await readFile(`${routesModFile}/${to.split('/')[0]}/routes.ts`, 'utf-8');

        if (separatorApp.search(to) < 0) {
            let app = to.search('/') < 0;
            if (app) {
                separatorApp = separatorApp.split('</script>');

                separatorApp.splice(separatorApp[0], 1, `${separatorApp[0]}import ${to} from '..${folderFront + to}/main.svelte';\n</script>`);
                separatorApp.splice(separatorApp[1].indexOf(), 0, `<${to} prefix='/${to}' />\n`);

                let finalRoute = separatorApp.join('');
                await writeFile(routesAppFile, finalRoute)
            } else {

                if (!(separatorMod.search(newTo) < 0)) {
                    throw new Error("YA EXISTE UN MODULO CON ESE NOMBRE")
                }

                separatorMod = separatorMod.split('const routes = {');

                separatorMod.splice(separatorMod[1], 1, `\n${separatorMod[0]}const routes = { '/${newTo[1]}': wrap({ asyncComponent: () => import('./${newTo[1]}/Modul.svelte') }),`);

                let finalRoute = separatorMod.join('');

                await writeFile(`${routesModFile}/${to.split('/')[0]}/routes.ts`, finalRoute)
            }
        }
    } catch (err) {
        // console.error(`Error al ejecutar el comando: ${err.message}`);
        throw new Error(err)
    }
}

async function existeApp(path) {
    try {
        console.log(`./FrontEnd${folderFront}` + path);
        return await stat(`./FrontEnd${folderFront}` + path);
    } catch (err) {
        return false;
    }
}

const tipo = process.argv[2];

switch (tipo) {
    case 'app':
        if (process.argv[3]) {
            copy('app/default', process.argv[3]);
            break;
        }
        console.log('Falta nombre de la aplicacion {nombre}');
        break;
    case 'mod':
        if (process.argv[3] && process.argv[4]) {
            existeApp(process.argv[3]).then(exist => {
                if (exist) {
                    copy(
                        'moduls/default',
                        process.argv[3] + '/' + process.argv[4],
                    );
                } else {
                    console.log('La aplicacion no existe');
                }
            });
            break;
        }
        console.log(
            'Falta nombre de la aplicacion o modulo {nombre app} {nombre modul}',
        );
        break;

    case 'help':
        help();
        break;

    default:
        console.log('Argumentos nesesarios: app | mod');
        break;
}

function help() {
    console.log(
        '----------------------------HELP-------------------------------',
    );
    console.log('npm run cli {comando principal} {argumento1} {argumento2}');
    console.log('');
    console.log('Comando principal:');
    console.log('-- app     Crea una aplicacion dentro de /app');
    console.log('           Ej: npm run cli {nombre aplicacion}');
    console.log('-- mod   Crea un modulo dentro de una aplicacion');
    console.log(
        '           Ej: npm run cli {nombre aplicacion} {monbre modulo}',
    );
    console.log('           la aplicaion debe estar cread previamente');
    console.log(
        '---------------------------------------------------------------',
    );
}
