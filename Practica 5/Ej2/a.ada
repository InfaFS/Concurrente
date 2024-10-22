--A
with Ada.Text_IO; use Ada.Text_IO;

procedure Ej2 is
    -- Definicion de la task Empleado
    Task Empleado is 
        Entry Pedido(D: IN texto; R:OUT texto);
    End empleado;

    --Definicion de cliente
    Task type Cliente;

    arrClientes: array(1..10) of Cliente;

    Task Body Cliente is
        Resultado: texto;
    begin
        Empleado.Pedido("datos",Resultado);
        Put_Line("Recibi la respuesta, me marcho");
    End Cliente;

    Task Body Empleado is
    begin
        loop
            accept Pedido(D: IN texto, R: OUT texto) do
                R := resolverPedido(D);
            end Pedido;
        end loop;
    End Empleado;

begin
    null;
end Ej2;