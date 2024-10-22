--A
with Ada.Text_IO; use Ada.Text_IO;

procedure Ej1 is
    -- Definición de la tarea del puente
    task Puente is
        Entry Pedido (Peso: IN Integer; R: OUT Boolean);
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
        while not paso loop
            Puente.Pedido(1, paso);  -- Peso de un Auto es 1
        end loop;
        if paso then
            Put_Line("Un Auto ha cruzado el puente.");
            Puente.Salir(1);
        end if;

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
                accept Pedido(Peso: IN Integer; R: OUT Boolean) do
                    if pesoActual + Peso <= 5 then
                        pesoActual := pesoActual + Peso;
                         Put_Line("Peso actual: " & Integer'Image(pesoActual));
                        R := True;  -- Permitir paso
                    else
                        R := False;  -- No permitir paso
                    end if;
                end Pedido;

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