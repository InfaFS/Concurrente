--5
with Ada.Text_IO; use Ada.Text_IO;

procedure Ej5 is
    -- Definicion de Jugador
    Task type Jugador is
        Entry EquipoGanador(IDEquipo: IN integer);
    end Jugador;

    --Definicion de Equipo
    Task type Equipo is
        Entry RecibirId(ID: IN integer);
        Entry Llegada();
        Entry Empezar();
        Entry Juntado(Suma: IN integer);
    End Equipo;

    --Definicion del admin que ve todo 
    Task Admin is
        Entry SumaEquipo(IDEquipo: IN integer; SumaEquipo: IN integer );
    End Admin;

    arrEquipos: array(1..5) of Equipo;

    arrJugadores: array(1..20) of Jugadores;

    --Jugador
    Task Body Jugador is
        IDEquipo: integer := ?;
        SumaMonedas: integer := 0;
    begin 
        arrEquipos(IDEquipo).Llegada(); --Le envian al equipo correspondiente que llegaron y espoeran a que les haya atendido
        arrEquipos(IDEquipo).Empezar(); --Esperan a que les atiendan el empezar
        for I in 1 .. 15 loop
            SumaMonedas := SumaMonedas + Moneda(); --Realiza la suma de las monedas
        end loop
        arrEquipos(IDEquipo).Juntado(SumaMonedas); --Le manda lo juntado al equipo 
        accept EquipoGanador(IDEquipo: IN integer); --Se queda esperando a que le manden el resultado;
    End Jugador;
    
    --Equipos
    Task Body Equipo is
        SumaTotalEquipo: integer :=0;
        IDEquipo : integer := -1;
    begin
        accept RecibirId(ID: IN integer) do --No va a saber su id asi que se lo manda el programa principal
            IDEquipo := ID;
        end RecibirId;

        for I in 1 .. 4 loop
            accept Llegada(); --Itera 4 veces para aceptar las llegdas
        end loop;

        for I in 1 .. 4 loop
            accept Empezar(); --Actuo lo anterior como barrera y ahora acepta los empezar de todos
        end loop;

        for I in 1 .. 4 loop
            accept Juntado(Suma: IN integer) do
                SumaTotalEquipo := SumaTotalEquipo + Suma; --Va aceptando las sumas y pone las sumas de todo
            end Juntado;
        end loop;

        Admin.SumaEquipo(IDEquipo,SumaTotalEquipo);  --Le manda al admin que ya tiene la suma junto con su id
    End Empleado;


    Task Body Admin is
        arrSumaEquipos: array (1..5) of integer := 0;
        sumaMaxima: integer := 0;
        EquipoGanador: integer := -1;
    begin
        for I in 1 .. 5 loop
            accept SumaEquipo(IDEquipo: IN integer,SumaEquipo: IN integer) do --Recibe la suma de los 5 equipos y los guarda en un array
                arrSumaEquipos(IDEquipo) := SumaEquipo;
            end Juntado;  
        end loop;   

        for I in 1 .. 5 loop
            if (arrSumaEquipos(I) > sumaMaxima) then --Se fija cual es el que tiene mas monedas
                EquipoGanador := I;
                sumaMaxima := arrSumaEquipos(I);
            end if
        end loop;  
        
        for I in 1..20 loop
            arrJugadores(I).EquipoGanador(EquipoGanador); --Le informa a cada usuario cual fue el equipo ganador
        end loop;

    End Admin;

begin
    for I in 1..5loop
        arrEquipos(I).RecibirId(I); --LE DA A CADA UNO EL ID
    end loop;
end Ej2;