--A
with Ada.Text_IO; use Ada.Text_IO;

procedure Ej1 is
    -- Definición de la tarea del puente
    task Puente is
        Entry PedidoA ();
        Entry PedidoC ();
        Entry PedidoCam ();

        Entry Salir (Peso: IN Integer);
    end Puente;

    -- Definición de las tareas de vehículos
    task type Auto;

    task type Camioneta;

    task type Camion;

    -- Arreglos de vehículos
    arrAutos: array(1..3) of Auto;
    arrCamionetas: array(1..3) of Camioneta;
    arrCamiones: array(1..3) of Camion;

    -- Implementación del cuerpo de las tareas de vehículos
    task body Auto is
        paso: Boolean := False;
    begin
        
        Puente.PedidoA();  -- Peso de un Auto es 1 modificar pada cada caso
        --Cruza        
        Puente.Salir(1);

    end Auto;

    task body Camioneta is
        paso: Boolean := False;
    begin
        while not paso loop
            Puente.Pedido(2, paso);  -- Peso de una Camioneta es 2
        end loop;
        
        if paso then
            Put_Line("Una Camioneta ha cruzado el puente.");
            Puente.Salir(2);
        end if;
  
    end Camioneta;

    task body Camion is
        paso: Boolean := False;
    begin
        while not paso loop
            Puente.Pedido(3, paso);  -- Peso de un Camión es 3
        end loop;
        if paso then
            Put_Line("Un Camión ha cruzado el puente.");
            Puente.Salir(3);
        end if;
     
    end Camion;

    -- Implementación del cuerpo de la tarea del puente
    task body Puente is
        pesoActual: Integer := 0;
    begin
        loop
            select
                when (pesoActual + 1) => 
                    accept PedidoA() do
                        pesoActual := pesoActual + 1;
                    end PedidoA;
                or 
                when (pesoActual + 2) => 
                    accept PedidoC() do
                        pesoActual := pesoActual + 2;
                    end PedidoC;
                or 
                when (pesoActual + 3) => 
                    accept PedidoCam() do
                        pesoActual := pesoActual + 3;
                    end PedidoC;
                or 
                accept Salir(Peso: IN Integer) do
                    pesoActual := pesoActual - Peso;  -- Reducir peso al salir
                end Salir;
            end select;
        end loop;
    end Puente;

begin
    Put_Line("Simulación del puente en marcha...");
end Ej1;

--b
--A
with Ada.Text_IO; use Ada.Text_IO;

procedure Ej1 is
    -- Definición de la tarea del puente
    task Puente is
        Entry PedidoA ();
        Entry PedidoC ();
        Entry PedidoCam ();

        Entry Salir (Peso: IN Integer);
    end Puente;

    -- Definición de las tareas de vehículos
    task type Auto;

    task type Camioneta;

    task type Camion;

    -- Arreglos de vehículos
    arrAutos: array(1..3) of Auto;
    arrCamionetas: array(1..3) of Camioneta;
    arrCamiones: array(1..3) of Camion;

    -- Implementación del cuerpo de las tareas de vehículos
    task body Auto is
        paso: Boolean := False;
    begin
        
        Puente.PedidoA();  -- Peso de un Auto es 1 modificar pada cada caso
        --Cruza        
        Puente.Salir(1);

    end Auto;

    task body Camioneta is
        paso: Boolean := False;
    begin
        while not paso loop
            Puente.Pedido(2, paso);  -- Peso de una Camioneta es 2
        end loop;
        
        if paso then
            Put_Line("Una Camioneta ha cruzado el puente.");
            Puente.Salir(2);
        end if;
  
    end Camioneta;

    task body Camion is
        paso: Boolean := False;
    begin
        while not paso loop
            Puente.Pedido(3, paso);  -- Peso de un Camión es 3
        end loop;
        if paso then
            Put_Line("Un Camión ha cruzado el puente.");
            Puente.Salir(3);
        end if;
     
    end Camion;

    -- Implementación del cuerpo de la tarea del puente
    task body Puente is
        pesoActual: Integer := 0;
    begin
        loop
            select
                when (pesoActual + 1) and (PedidoCam'count = 0) => 
                    accept PedidoA() do
                        pesoActual := pesoActual + 1;
                    end PedidoA;
                or 
                when (pesoActual + 2) and (PedidoCam'count = 0) => 
                    accept PedidoC() do
                        pesoActual := pesoActual + 2;
                    end PedidoC;
                or 
                when (pesoActual + 3) => 
                    accept PedidoCam() do
                        pesoActual := pesoActual + 3;
                    end PedidoC;
                or 
                accept Salir(Peso: IN Integer) do
                    pesoActual := pesoActual - Peso;  -- Reducir peso al salir
                end Salir;
            end select;
        end loop;
    end Puente;

begin
    Put_Line("Simulación del puente en marcha...");
end Ej1;