--D
with Ada.Text_IO; use Ada.Text_IO;

procedure Ej2 is
    -- Definicion de la task Empleado
    Task Empleado is 
        Entry Pedido(D: IN String; R:OUT String);
    End empleado;

    --Definicion de cliente
    Task type Cliente;

    arrClientes: array(1..10) of Cliente;

    Task Body Cliente is
        Resultado: string;
    begin
        SELECT
            Empleado.Pedido("datos",Resultado);
            Put_Line("Recibi la respuesta, me marcho");
        OR DELAY 600.0;
            Put_Line("No me atendieron en 10 min, pruebo una vez mas");
            SELECT 
                Empleado.Pedido("datos",Resultado);
                Put_Line("Recibi la respuesta, me marcho");
            OR DELAY 0.0;
                Put_Line("Me marcho, no me atendieron por segunda vez");
            END SELECT;
        END SELECT;
    End Cliente;

    Task Body Empleado is
    begin
        loop
            accept Pedido(D: IN String, R: OUT String) do
                --R := resolverPedido(D);
                R:= "Respuesta enviada";
            end Pedido;
        end loop;
    End Empleado;

begin
    null;
end Ej2;