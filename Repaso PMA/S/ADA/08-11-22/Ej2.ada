Procedure Ej3 is 

-----
Task Empleado is
    Entry pedidoRegular(Pedido: IN texto, Comprobante: OUT texto);
    Entry pedidoPremium(Pedido: IN texto, Comprobante: OUT texto);
End Empleado;

-----
Task type ClienteRegular;


-----
Task type ClientePremium;


arrPremium: array (1..P) of ClientePremium;
arrRegular: array (1..R) of ClienteRegular;

Task Body Empleado is

Begin
    loop
        SELECT
            when(pedidoPremium'count = 0) =>
                accept pedidoRegular(Pedido: IN texto, Comprobante: OUT texto) do
                    Comprobante = realizarPago(Pedido); 
                end pedidoRegular;
        OR
            accept pedidoPremium(Pedido: IN texto, Comprobante: OUT texto) do
                Comprobante = realizarPago(Pedido); 
            end pedidoPremium;
        END SELECT;
    end loop;
End Empleado;


Task Body ClientePremium is
    Pedido: pedido;
Begin
    pedido := generarPedido();
    Empleado.pedidoPremium(pedido);

End ClientePremium;


Task Body ClienteRegular is
    Pedido: pedido;
Begin

    pedido := generarPedido();
    SELECT
        Empleado.pedidoRegular(pedido);
    OR DELAY 1800;
        null;
        -- o podemos poner un put_line("me voy");
    END SELECT;
End ClienteRegular;




Begin
    null;
End Ej3;