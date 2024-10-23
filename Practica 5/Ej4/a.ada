--Ej 4
with Ada.Text_IO; use Ada.Text_IO;

procedure Ej4 is
    -- Definicion de la task Empleado
    Task Medico is 
        Entry PedidoEnfermo(D: IN Solicitud; R:OUT Diagnostico);
    End Medico;

    --Definicion de cliente
    Task type Persona is 
    --
    End Persona;

    arrEnfermos: array(1..P) of Persona;

    Task type Enfermera is
    --
    End Enfermera;

    arrEnfermeras: array(1..E) of Enfermera;


    Task Body Persona is
        Solicitud: Solicitud;
        Diagnostico: Diagnostico;
    begin
        SELECT
            Medico.PedidoEnfermo(Solicitud,Diagnostico);
            Put_Line("Recibi el diagnostico, me voy");
        OR DELAY 300.0;
            Put_Line("No me atendieron en 5 min");
            SELECT 
                Medico.PedidoEnfermo(Solicitud,Diagnostico);
                Put_Line("Recibi el diagnostico, me voy");
            OR DELAY 600.0;
                Put_Line("No me atendieron en 10 min ,pruebo una vez mas");
                SELECT
                    Medico.PedidoEnfermo(Solicitud,Diagnostico);
                    Put_Line("Recibi el diagnostico, me voy");
                OR DELAY 0.0;
                    Put_Line("Me voy!");
                END SELECT;
            END SELECT;
        END SELECT;
    End Persona;

    --- Task medico
    Task Body Medico is

    begin
        loop
            SELECT
            accept PedidoEnfermo(S: IN Solicitud, R: OUT Diagnostico) do
                R := resolverEnfermedad(S);
            end PedidoEnfermo;
            OR
                when(PedidoEnfermo'count = 0) =>
                    accept PedidoEnfermera(S: IN Solicitud);
            OR
                when(PedidoEnfermo'Count = 0 and PedidoEnfermera = 0) => 
                    accept PedidosEscritorio(S: IN Solicitud);
            END SELECT;
        end loop;
    End Medico;

    --Admin escritorio

    Task Body AdminEscritorio is
    arrSolicitudes: array of Solicitud;
    begin
        loop
            SELECT
                accept NotaEscritorio(S: IN Solicitud) do
                    arrSolicitudes.push(S);
                end NotaEscritorio;
            OR
                when(NotaEscritorio'count = 0) =>
                    Medico.PedidosEscritorio(arrSolicitudes.pop());
            END SELECT;
        end loop;
    End AdminEscritorio;

    --Enfermera
    Task Body Enfermera is
        Solicitud: Solicitud;
        loop
            SELECT
                medico.PedidoEnfermera(Solicitud);
            OR DELAY 0.0;
                AdminEscritorio.NotaEscritorio(Solicitud);
            END SELECT;
        end loop;
    End Enfermera;

begin
    null;
end Ej4;

--El medico va a ser como el admin el cual va a estar recibiendo tanto de enfermeras como de pacientes, los papeles los vamos a guardar en una task tipo admin pero de buffers
