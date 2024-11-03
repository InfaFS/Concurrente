Procedure Ej3 is

Task type Controlador;


arrControladores: array (1..150) of Controlador;

Task ModuloCentral is
    entry PedidoSuperaRango(Temperatura: IN real; Accion: OUT string);
    entry PedidoDebajoRango(Temperatura: IN real; Accion: OUT string);
end ModuloCentral;

Task body ModuloCentral is
Begin
    loop
        SELECT
            accept PedidoSuperaRango(Temperatura: IN real; Accion: OUT string)do
                Accion = determinar(Temperatura);
            end PedidoSuperaRango;
        OR
            when (PedidoSuperaRango'count = 0) =>
                accept PedidoDebajoRango(Temperatura: IN real; Accion: OUT string)do
                    Accion = determinar(Temperatura);
                end PedidoDebajoRango;
        END SELECT;
    end loop;
end ModuloCentral;

Task Body Controlador is
    Temperatura: int;
    Accion: string;
Begin
    loop
        Temperatura := medir();
        if(Temperatura > rango) then
            ModuloCentral.PedidoSuperaRango(Temperatura,Accion);
            actualizar(Accion);
        elsif then
            SELECT
                ModuloCentral.PedidoDebajoRango(Temperatura,Accion);
                actualizar(Accion);
            OR DELAY(600.0);
                null;
            END SELECT;
        end if;
        delay (30.0);
    end loop;
End Controlador;




Begin
    null;
End Ej3;