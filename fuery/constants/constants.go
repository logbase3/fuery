/*
	Fuery, is a small and simple tool for querying files using SQL.
	Copyright (C) 2013 log₃() <contact@logbase3.com>

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with this program. If not, see <http://www.gnu.org/licenses/>.

	For more information visit https://github.com/logbase3/fuery/fuery
	or send an e-mail to contact@logbase3.com
*/

// Package constants only exists to clarify when I'm using a constant in the code.
package constants // import "logbase3.com/fuery/fuery/constants"

const (
	// This constant is just the name of the file and should not be used for
	// opening the file. Instead the HISTORY_FILE variable should be used.
	HistoryFileName string = "/.fuery_history"

	// Current program version.
	Version string = "0.1"

	// Current year
	Year string = "2015"

	// Program header
	Header string = "    Fuery v" + Version + " Copyright (C) " + Year + " log₃()" + `
    This program comes with ABSOLUTELY NO WARRANTY.
    This is free software, and you are welcome to redistribute it
    under certain conditions. Read the license file or visit:
    'http://www.gnu.org/copyleft/gpl.html' for more information.

    Use '\h' to get help about available options or '\?' to list
    system commands. For more information please visit:
    http://github.com/logbase3/fuery

`

	// License text
	License string = `    Fuery, is a small and simple tool for querying files using SQL.
    Copyright (C) 2013 log₃() <contact@logbase3.com>

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or any
    later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this program. If not, see <http://www.gnu.org/licenses/>.

    For more information visit https://github.com/logbase3/fuery
    or send an e-mail to contact@logbase3.com`

	SystemHelp string = `General
  \copyright      mostrar tÚrminos de uso y distribuci¾n de PostgreSQL
  \g [ARCH] o ;   enviar b·fer de consulta al servidor (y resultados a archivo u |orden)
  \gset [PREFIJO] ejecutar la consulta y almacenar los resultados en variables de psql
  \h [NOMBRE]     mostrar ayuda de sintaxis de ¾rdenes SQL; use ½*╗ para todas las ¾rdenes
  \q              salir de psql
  \watch [SEGS]   ejecutar consulta cada SEGS segundos

B·fer de consulta
  \e [ARCHIVO] [L═NEA] editar el b·fer de consulta (o archivo) con editor externo
  \ef [NOMBRE-FUNCIËN [L═NEA]] editar una funci¾n con editor externo
  \p              mostrar el contenido del b·fer de consulta
  \r              reiniciar (limpiar) el b·fer de consulta
  \w ARCHIVO      escribir b·fer de consulta a archivo

Entrada/Salida
  \copy ...       ejecutar orden SQL COPY con flujo de datos al cliente
  \echo [CADENA]  escribir cadena a salida estßndar
  \i ARCHIVO      ejecutar ¾rdenes desde archivo
  \ir ARCHIVO     como \i, pero relativo a la ubicaci¾n del script actual
  \o [ARCHIVO]    enviar resultados de consultas a archivo u |orden
  \qecho [CADENA] escribir cadena a salida de consultas (ver \o)

Informativo
   (opciones: S = desplegar objectos de sistema, + = agregar mßs detalle)
  \d[S+]            listar tablas, vistas y secuencias
  \d[S+]  NOMBRE    describir tabla, Ýndice, secuencia o vista
  \da[S]  [PATRËN]  listar funciones de agregaci¾n
  \db[+]  [PATRËN]  listar tablespaces
  \dc[S+] [PATRËN]  listar conversiones
  \dC[+]  [PATRËN]  listar conversiones de tipo (casts)
  \dd[S]  [PATRËN]  listar comentarios de objetos que no aparecen en otra parte
  \ddp    [PATRËN]  listar privilegios por omisi¾n
  \dD[S+] [PATRËN]  listar dominios
  \det[+] [PATRËN]  listar tablas forßneas
  \des[+] [PATRËN]  listar servidores forßneos
  \deu[+] [PATRËN]  listar mapeos de usuario
  \dew[+] [PATRËN]  listar conectores de datos externos
  \df[antw][S+] [PATRËN]  listar funciones [s¾lo ag./normal/trigger/ventana]
  \dF[+]  [PATRËN]  listar configuraciones de b·squeda en texto
  \dFd[+] [PATRËN]  listar diccionarios de b·squeda en texto
  \dFp[+] [PATRËN]  listar analizadores (parsers) de b·sq. en texto
  \dFt[+] [PATRËN]  listar plantillas de b·squeda en texto
  \dg[+]  [PATRËN]  listar roles
  \di[S+] [PATRËN]  listar Ýndices
  \dl               listar objetos grandes, lo mismo que \lo_list
  \dL[S+] [PATRËN]  listar lenguajes procedurales
  \dm[S+] [PATRËN]  listar vistas materializadas
  \dn[S+] [PATRËN]  listar esquemas
  \do[S]  [PATRËN]  listar operadores
  \dO[S]  [PATRËN]  listar ordenamientos (collations)
  \dp     [PATRËN]  listar privilegios de acceso a tablas, vistas y secuencias
  \drds [PAT1 [PAT2]] listar parßmetros de rol por base de datos
  \ds[S+] [PATRËN]  listar secuencias
  \dt[S+] [PATRËN]  listar tablas
  \dT[S+] [PATRËN]  listar tipos de dato
  \du[+]  [PATRËN]  listar roles
  \dv[S+] [PATRËN]  listar vistas
  \dE[S+] [PATRËN]  listar tablas forßneas
  \dx[+]  [PATRËN]  listar extensiones
  \dy    [PATRËN]   listar disparadores por eventos
  \l[+] [PATRËN]    listar bases de datos
  \sf[+] FUNCIËN    mostrar la definici¾n de una funci¾n
  \z      [PATRËN]  lo mismo que \dp

Formato
  \a              cambiar entre modo de salida alineado y sin alinear
  \C [CADENA]     definir tÝtulo de tabla, o indefinir si es vacÝo
  \f [CADENA]     mostrar o definir separador de campos para modo de salida sin alinear
  \H              cambiar modo de salida HTML (actualmente desactivado)
  \pset NOMBRE [VALOR] define opci¾n de salida de tabla (NOMBRE = format,border, expanded,fieldsep,fieldsep_zero,footer,null,numericlocale, recordsep,recordsep_zero,tuples_only,title,tableattr,pager)
  \t [on|off]     mostrar s¾lo filas (actualmente desactivado)
  \T [CADENA]     definir atributos HTML de <table>, o indefinir si es vacÝo
  \x [on|off|auto] cambiar modo expandido (actualmente desactivado)

Conexiones
  \c[onnect] [BASE-DE-DATOS|- USUARIO|- ANFITRIËN|- PUERTO|-] conectar a una nueva base de datos (actual: ½postgres╗)
  \encoding [CODIFICACIËN] mostrar o definir codificaci¾n del cliente
  \password [USUARIO] cambiar la contrase±a para un usuario en forma segura
  \conninfo       despliega la informaci¾n sobre la conexi¾n actual

Sistema Operativo
  \cd [DIR]        cambiar el directorio de trabajo actual
  \setenv NOMBRE [VALOR] definir o indefinir variable de ambiente
  \timing [on|off] mostrar tiempo de ejecuci¾n de ¾rdenes (actualmente desactivado)
  \! [ORDEN]       ejecutar orden en intÚrprete de ¾rdenes (shell), o iniciar intÚrprete interactivo

Variables
  \prompt [TEXTO] NOMBRE  preguntar al usuario el valor de la variable
  \set [NOMBRE [VALOR]]   definir variables internas, listar todas si no se dan parßmetros
  \unset NOMBRE           indefinir (eliminar) variable interna

Objetos Grandes
  \lo_export LOBOID ARCHIVO
  \lo_import ARCHIVO [COMENTARIO]
  \lo_list
  \lo_unlink LOBOID   operaciones con objetos grandes`
)
