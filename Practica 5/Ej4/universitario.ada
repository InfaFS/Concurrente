--Universitarios
with Ada.Text_IO; use Ada.Text_IO;

procedure EjUniversitario is
    -- Definicion de la task Servidor
    Task Servidor is 
        Entry Pedido(D: IN texto; R:OUT String);
    End empleado;

    --Definicion de Usuario
    Task type Usuario;

    arrUsuarios: array(1..10) of Usuario;

    Task Body Usuario is
        Documento: texto;
        Respuesta: String := "Contiene Error";
        RespuestaServer: bool := false;
    begin
        --Una vez que tengo el documentom, entro al select para transferirle al server
        while Respuesta = "Contiene Error" loop
            --Trabajo el documento
            TrabajarDocumento(Documento);
            --Loop que itera hasta que el server me diga una respuesta
            while not respuestaServer loop
                SELECT
                    Servidor.Pedido(Documento,Respuesta);
                    respuestaServer := true;
                OR DELAY 120.0;
                    DELAY 60.0;
                END SELECT;
            end loop;
            respuestaServer := false;
            --si la respuesta fue Contiene Error entonces repite, sino finaliza
        end loop;
    End Usuario;

    Task Body Server is
    Aprobado: bool := false;
    begin
        loop
            accept Pedido(D: IN texto, R: OUT String) do
                Aprobado := resolverPedido(D);
                if (not Aprobado) then
                    R := "Contiene Error";
                end if
                else
                    R := "Aprobado!";
                end if
            end Pedido;
        end loop;
    End Server;

begin
    null;
end EjUniversitario;